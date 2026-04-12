const API_URL = "http://localhost:8088"

// guarda o valor real para mostrar/ocultar
let paymentsTotal = null
let paymentsVisible = true

// verifica se o admin está autenticado, senão redireciona pro login
function checkAuth() {
    const token = localStorage.getItem("token")
    if (!token) {
        window.location.href = "../index.html"
    }
    return token
}

// preenche o nome do admin na sidebar
function loadAdminName() {
    const admin = JSON.parse(localStorage.getItem("admin"))
    if (admin) {
        document.getElementById("adminName").textContent = admin.name
    }
}

// busca os associados ativos na API e preenche a tabela e o card
async function loadAssociates() {
    const token = checkAuth()

    try {
        const response = await fetch(`${API_URL}/associates`, {
            method: "GET",
            headers: {
                "Authorization": `Bearer ${token}`
            }
        })

        if (response.status === 401) {
            logout()
            return
        }

        const associates = await response.json()

        document.getElementById("totalAssociates").textContent = associates.length

        const tbody = document.getElementById("associatesTable")

        if (associates.length === 0) {
            tbody.innerHTML = `
                <tr>
                    <td colspan="4" class="table-loading">Nenhum associado ativo encontrado.</td>
                </tr>
            `
            return
        }
        const preview = associates.slice(0, 10)


        tbody.innerHTML = associates.map(associate => `
            <tr>
                <td>${associate.name}</td>
                <td>${formatCPF(associate.cpf)}</td>
                <td>${associate.position}</td>
                <td>
                    <span class="badge ${associate.status ? 'badge-active' : 'badge-inactive'}">
                        ${associate.status ? 'Ativo' : 'Inativo'}
                    </span>
                </td>
            </tr>
        `).join("")

    } catch (error) {
        const tbody = document.getElementById("associatesTable")
        tbody.innerHTML = `
            <tr>
                <td colspan="4" class="table-loading">Erro ao carregar associados.</td>
            </tr>
        `
    }
}

// busca o total de pagamentos do mês atual
async function loadMonthlyTotal() {
    const token = checkAuth()

    const today = new Date()
    const competence = `${today.getFullYear()}-${String(today.getMonth() + 1).padStart(2, "0")}-01`

    try {
        const response = await fetch(`${API_URL}/payments/month?competence=${competence}`, {
            method: "GET",
            headers: {
                "Authorization": `Bearer ${token}`
            }
        })

        if (response.status === 401) {
            logout()
            return
        }

        const data = await response.json()

        // salva o valor para usar no toggle
        paymentsTotal = data.total
        document.getElementById("totalPayments").textContent = formatCurrency(paymentsTotal)

    } catch (error) {
        document.getElementById("totalPayments").textContent = "—"
    }
}

// alterna visibilidade do valor de pagamentos
function togglePayments() {
    const el = document.getElementById("totalPayments")
    const btn = document.getElementById("btnToggle")

    paymentsVisible = !paymentsVisible

    if (paymentsVisible) {
        el.textContent = paymentsTotal !== null ? formatCurrency(paymentsTotal) : "—"
        btn.textContent = "👁️"
        btn.title = "Ocultar valor"
    } else {
        el.innerHTML = "<span class='payment-hidden'></span>"
        btn.textContent = "👁️"  // ← mesma coisa
        btn.title = "Mostrar valor"
    }
}

// formata CPF: 00000000000 → 000.000.000-00
function formatCPF(cpf) {
    return cpf.replace(/(\d{3})(\d{3})(\d{3})(\d{2})/, "$1.$2.$3-$4")
}

// formata valor em reais: 1500 → R$ 1.500,00
function formatCurrency(value) {
    return value.toLocaleString("pt-BR", {
        style: "currency",
        currency: "BRL"
    })
}

// desloga o admin limpando o localStorage e voltando pro login
function logout() {
    localStorage.removeItem("token")
    localStorage.removeItem("admin")
    window.location.href = "../index.html"
}

// inicializa o dashboard
loadAdminName()
loadAssociates()
loadMonthlyTotal()