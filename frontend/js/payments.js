const API_URL = "http://localhost:8088"

// verifica autenticação
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

// inicializa o seletor com o mês atual
function initMonthPicker() {
    const today = new Date()
    const year = today.getFullYear()
    const month = String(today.getMonth() + 1).padStart(2, "0")
    document.getElementById("monthPicker").value = `${year}-${month}`
}

// retorna a competence no formato yyyy-mm-dd a partir do input[type=month]
function getCompetence() {
    const value = document.getElementById("monthPicker").value
    if (!value) return null
    return `${value}-01`
}

// carrega pagamentos e inadimplentes ao mesmo tempo
async function loadAll() {
    await Promise.all([loadPayments(), loadDefaulters()])
}

// busca os pagamentos do mês selecionado
async function loadPayments() {
    const token = checkAuth()
    const competence = getCompetence()
    if (!competence) return

    const tbody = document.getElementById("paymentsTable")
    tbody.innerHTML = `<tr><td colspan="5" class="table-loading">Carregando...</td></tr>`
    document.getElementById("totalPayments").textContent = "—"
    document.getElementById("totalCount").textContent = "—"

    try {
        const response = await fetch(`${API_URL}/payments?competence=${competence}`, {
            method: "GET",
            headers: { "Authorization": `Bearer ${token}` }
        })

        if (response.status === 401) { logout(); return }

        const payments = await response.json()
        renderPaymentsTable(payments)
        renderSummary(payments)

    } catch (error) {
        tbody.innerHTML = `<tr><td colspan="5" class="table-loading">Erro ao carregar pagamentos.</td></tr>`
    }
}

// busca os inadimplentes do mês selecionado
async function loadDefaulters() {
    const token = checkAuth()
    const competence = getCompetence()
    if (!competence) return

    const tbody = document.getElementById("defaultersTable")
    tbody.innerHTML = `<tr><td colspan="2" class="table-loading">Carregando...</td></tr>`
    document.getElementById("totalDefaulters").textContent = "—"

    try {
        const response = await fetch(`${API_URL}/payments/defaulters?competence=${competence}`, {
            method: "GET",
            headers: { "Authorization": `Bearer ${token}` }
        })

        if (response.status === 401) { logout(); return }

        const defaulters = await response.json()
        renderDefaultersTable(defaulters)
        document.getElementById("totalDefaulters").textContent = defaulters.length

    } catch (error) {
        tbody.innerHTML = `<tr><td colspan="2" class="table-loading">Erro ao carregar inadimplentes.</td></tr>`
    }
}

// renderiza a tabela de pagamentos
function renderPaymentsTable(payments) {
    const tbody = document.getElementById("paymentsTable")

    if (payments.length === 0) {
        tbody.innerHTML = `<tr><td colspan="5" class="table-loading">Nenhum pagamento encontrado para este mês.</td></tr>`
        return
    }

    tbody.innerHTML = payments.map(p => `
        <tr>
            <td>${p.associate_name}</td>
            <td>${formatDate(p.competence)}</td>
            <td>${formatDate(p.payment_date)}</td>
            <td>${formatCurrency(p.value)}</td>
            <td>
                <span class="badge ${p.status ? 'badge-paid' : 'badge-pending'}">
                    ${p.status ? 'Pago' : 'Pendente'}
                </span>
            </td>
        </tr>
    `).join("")
}

// renderiza a tabela de inadimplentes
function renderDefaultersTable(defaulters) {
    const tbody = document.getElementById("defaultersTable")

    if (defaulters.length === 0) {
        tbody.innerHTML = `<tr><td colspan="2" class="table-loading">Nenhum inadimplente neste mês. 🎉</td></tr>`
        return
    }

    tbody.innerHTML = defaulters.map(d => `
        <tr>
            <td><a href="associate_detail.html?id=${d.id}" class="associate-link">${d.name}</a></td>
            <td>${formatCPF(d.cpf)}</td>
        </tr>
    `).join("")
}

// atualiza os cards de resumo
function renderSummary(payments) {
    const total = payments.reduce((acc, p) => acc + p.value, 0)
    document.getElementById("totalPayments").textContent = formatCurrency(total)
    document.getElementById("totalCount").textContent = payments.length
}

// formata data yyyy-mm-dd para dd/mm/yyyy
function formatDate(dateStr) {
    if (!dateStr) return "—"
    const [year, month, day] = dateStr.split("T")[0].split("-")
    return `${day}/${month}/${year}`
}

// formata valor em reais
function formatCurrency(value) {
    return value.toLocaleString("pt-BR", {
        style: "currency",
        currency: "BRL"
    })
}

// formata CPF: 00000000000 → 000.000.000-00
function formatCPF(cpf) {
    return cpf.replace(/(\d{3})(\d{3})(\d{3})(\d{2})/, "$1.$2.$3-$4")
}

// logout
function logout() {
    localStorage.removeItem("token")
    localStorage.removeItem("admin")
    window.location.href = "../index.html"
}

// inicializa
loadAdminName()
initMonthPicker()
loadAll()