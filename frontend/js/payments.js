const API_URL = "http://localhost:8088"

// lista de associados ativos carregada uma vez ao abrir o modal
let allAssociates = []

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
            <td><a href="associate_detail.html?id=${p.associate_id}" class="associate-link">${p.associate_name}</a></td>
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

// abre modal e carrega lista de associados
async function openModal() {
    const token = checkAuth()
    const errorEl = document.getElementById("modalError")
    errorEl.textContent = ""

    // preenche a competência com o mês selecionado
    const monthValue = document.getElementById("monthPicker").value
    document.getElementById("inputCompetence").value = monthValue

    // preenche a data de pagamento com hoje
    const today = new Date().toISOString().split("T")[0]
    document.getElementById("inputPaymentDate").value = today
    document.getElementById("inputPaymentDate").max = today

    document.getElementById("inputValue").value = ""
    document.getElementById("inputAssociate").value = ""
    document.getElementById("inputAssociateID").value = ""
    document.getElementById("autocompleteList").classList.remove("active")

    // carrega associados ativos
    try {
        const response = await fetch(`${API_URL}/associates`, {
            method: "GET",
            headers: { "Authorization": `Bearer ${token}` }
        })

        if (response.status === 401) { logout(); return }

        allAssociates = await response.json()

    } catch (error) {
        errorEl.textContent = "Erro ao carregar associados."
    }

    document.getElementById("modalOverlay").classList.add("active")
}

// filtra associados conforme o admin digita
function filterAssociates() {
    const query = document.getElementById("inputAssociate").value.toLowerCase().trim()
    const list = document.getElementById("autocompleteList")

    // limpa o ID selecionado quando o admin edita o campo
    document.getElementById("inputAssociateID").value = ""

    if (query === "") {
        list.classList.remove("active")
        list.innerHTML = ""
        return
    }

    const filtered = allAssociates.filter(a =>
        a.name.toLowerCase().includes(query)
    )

    if (filtered.length === 0) {
        list.classList.remove("active")
        return
    }

    list.innerHTML = filtered.map(a => `
        <li onclick="selectAssociate(${a.id}, '${a.name}')">${a.name}</li>
    `).join("")

    list.classList.add("active")
}

// preenche o campo ao selecionar uma sugestão
function selectAssociate(id, name) {
    document.getElementById("inputAssociate").value = name
    document.getElementById("inputAssociateID").value = id
    document.getElementById("autocompleteList").classList.remove("active")
    document.getElementById("autocompleteList").innerHTML = ""
}

// fecha modal
function closeModal() {
    document.getElementById("modalOverlay").classList.remove("active")
    document.getElementById("autocompleteList").classList.remove("active")
}

// fecha modal ao clicar fora
function closeModalOnOverlay(event) {
    if (event.target === document.getElementById("modalOverlay")) {
        closeModal()
    }
}

// registra o pagamento
async function savePayment() {
    const token = checkAuth()
    const errorEl = document.getElementById("modalError")
    const btnSave = document.getElementById("btnSave")

    errorEl.textContent = ""

    const associateID = parseInt(document.getElementById("inputAssociateID").value)
    const competence = document.getElementById("inputCompetence").value
    const paymentDate = document.getElementById("inputPaymentDate").value
    const value = parseFloat(document.getElementById("inputValue").value)

    // validações
    if (!associateID) {
        errorEl.textContent = "Selecione um associado da lista."
        return
    }

    if (!competence) {
        errorEl.textContent = "Informe a competência."
        return
    }

    if (!paymentDate) {
        errorEl.textContent = "Informe a data do pagamento."
        return
    }

    if (!value || value <= 0) {
        errorEl.textContent = "Informe um valor válido."
        return
    }

    btnSave.disabled = true
    btnSave.textContent = "Registrando..."

    try {
        const response = await fetch(`${API_URL}/payment`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                "Authorization": `Bearer ${token}`
            },
            body: JSON.stringify({
                associateID,
                competence: `${competence}-01`,
                paymentDate,
                value,
                status: true
            })
        })

        const data = await response.json()

        if (!response.ok) {
            errorEl.textContent = data.error || "Erro ao registrar pagamento."
            return
        }

        closeModal()
        loadAll()

    } catch (error) {
        errorEl.textContent = "Não foi possível conectar ao servidor."
    } finally {
        btnSave.disabled = false
        btnSave.textContent = "Registrar"
    }
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

// fecha autocomplete ao clicar fora dele
document.addEventListener("click", function(e) {
    const wrapper = document.querySelector(".autocomplete-wrapper")
    if (wrapper && !wrapper.contains(e.target)) {
        document.getElementById("autocompleteList")?.classList.remove("active")
    }
})

document.addEventListener("keydown", function(e) {
    const modalActive = document.getElementById("modalOverlay").classList.contains("active")
    if (e.key === "Enter" && modalActive) {
        savePayment()
    }
    if (e.key === "Escape" && modalActive) {
        closeModal()
    }
})

// inicializa
loadAdminName()
initMonthPicker()
loadAll()