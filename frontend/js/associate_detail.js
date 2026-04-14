const API_URL = "http://localhost:8088"

let currentAssociate = null

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

// pega o id da URL: associate_detail.html?id=12
function getIDFromURL() {
    const params = new URLSearchParams(window.location.search)
    return params.get("id")
}

// busca o associado pelo ID e preenche a ficha
async function loadAssociate() {
    const token = checkAuth()
    const id = getIDFromURL()

    if (!id) {
        window.location.href = "associates.html"
        return
    }

    try {
        const response = await fetch(`${API_URL}/associate/id/${id}`, {
            method: "GET",
            headers: { "Authorization": `Bearer ${token}` }
        })

        if (response.status === 401) { logout(); return }

        if (!response.ok) {
            window.location.href = "associates.html"
            return
        }

        currentAssociate = await response.json()
        renderDetail(currentAssociate)

    } catch (error) {
        window.location.href = "associates.html"
    }
}

// preenche a ficha com os dados do associado
function renderDetail(a) {
    const initials = a.name.split(" ").map(n => n[0]).slice(0, 2).join("").toUpperCase()
    document.getElementById("detailAvatar").textContent = initials

    document.getElementById("detailName").textContent = a.name

    const badge = document.getElementById("detailBadge")
    badge.textContent = a.status ? "Ativo" : "Inativo"
    badge.className = `badge ${a.status ? "badge-active" : "badge-inactive"}`

    document.getElementById("detailCPF").textContent = formatCPF(a.cpf)
    document.getElementById("detailEmail").textContent = a.email
    document.getElementById("detailTel").textContent = formatTel(a.tel)
    document.getElementById("detailPosition").textContent = a.position
    document.getElementById("detailDateOfBirth").textContent = formatDate(a.date_of_birth)
    document.getElementById("detailAssociationDate").textContent = formatDate(a.association_date)
    document.getElementById("detailAddress").textContent = a.address

    document.title = `CESJB — ${a.name}`
}

// abre modal de edição preenchido
function openEditModal() {
    if (!currentAssociate) return

    document.getElementById("associateId").value = currentAssociate.id
    document.getElementById("inputName").value = currentAssociate.name
    document.getElementById("inputCPF").value = formatCPF(currentAssociate.cpf)
    document.getElementById("inputEmail").value = currentAssociate.email
    document.getElementById("inputTel").value = currentAssociate.tel
    document.getElementById("inputDateOfBirth").value = currentAssociate.date_of_birth?.split("T")[0] || ""
    document.getElementById("inputAssociationDate").value = currentAssociate.association_date?.split("T")[0] || ""
    document.getElementById("inputPosition").value = currentAssociate.position
    document.getElementById("inputAddress").value = currentAssociate.address
    document.getElementById("inputStatus").value = currentAssociate.status ? "true" : "false"
    document.getElementById("modalError").textContent = ""
    document.getElementById("modalOverlay").classList.add("active")
}

// fecha modal
function closeModal() {
    document.getElementById("modalOverlay").classList.remove("active")
}

// fecha modal ao clicar fora
function closeModalOnOverlay(event) {
    if (event.target === document.getElementById("modalOverlay")) {
        closeModal()
    }
}

// salva edição do associado
async function saveAssociate() {
    const token = checkAuth()
    const id = document.getElementById("associateId").value
    const errorEl = document.getElementById("modalError")
    const btnSave = document.getElementById("btnSave")

    errorEl.textContent = ""

    const dateOfBirth = document.getElementById("inputDateOfBirth").value
    const associationDate = document.getElementById("inputAssociationDate").value

    if (!dateOfBirth) {
        errorEl.textContent = "Data de nascimento é obrigatória."
        return
    }

    if (!associationDate) {
        errorEl.textContent = "Data de associação é obrigatória."
        return
    }

    const body = {
        name: toTitleCase(document.getElementById("inputName").value.trim()),
        cpf: document.getElementById("inputCPF").value.replace(/\D/g, ""),
        email: document.getElementById("inputEmail").value.trim(),
        tel: document.getElementById("inputTel").value.replace(/\D/g, ""),
        date_of_birth: dateOfBirth,
        association_date: associationDate,
        position: document.getElementById("inputPosition").value.trim(),
        address: document.getElementById("inputAddress").value.trim(),
        status: document.getElementById("inputStatus").value === "true"
    }

    if (!body.name || !body.cpf || !body.email) {
        errorEl.textContent = "Nome, CPF e e-mail são obrigatórios."
        return
    }

    if (!isValidEmail(body.email)) {
        errorEl.textContent = "O e-mail informado não é válido."
        return
    }

    btnSave.disabled = true
    btnSave.textContent = "Salvando..."

    try {
        const response = await fetch(`${API_URL}/associate/${id}`, {
            method: "PUT",
            headers: {
                "Content-Type": "application/json",
                "Authorization": `Bearer ${token}`
            },
            body: JSON.stringify(body)
        })

        const data = await response.json()

        if (!response.ok) {
            errorEl.textContent = data.error || "Erro ao salvar associado."
            return
        }

        closeModal()
        loadAssociate()

    } catch (error) {
        errorEl.textContent = "Não foi possível conectar ao servidor."
    } finally {
        btnSave.disabled = false
        btnSave.textContent = "Salvar"
    }
}

// desativa associado
async function deactivateAssociate() {
    if (!confirm("Deseja desativar este associado?")) return

    const token = checkAuth()

    try {
        const body = {
            ...currentAssociate,
            cpf: currentAssociate.cpf.replace(/\D/g, ""),
            date_of_birth: currentAssociate.date_of_birth?.split("T")[0],
            association_date: currentAssociate.association_date?.split("T")[0],
            status: false
        }

        const response = await fetch(`${API_URL}/associate/${currentAssociate.id}`, {
            method: "PUT",
            headers: {
                "Content-Type": "application/json",
                "Authorization": `Bearer ${token}`
            },
            body: JSON.stringify(body)
        })

        if (!response.ok) {
            alert("Erro ao desativar associado.")
            return
        }

        window.location.href = "associates.html"

    } catch (error) {
        alert("Não foi possível conectar ao servidor.")
    }
}

// valida formato de email
function isValidEmail(email) {
    return /^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(email)
}

// converte string para Title Case: "joao silva" → "Joao Silva"
function toTitleCase(str) {
    return str.toLowerCase().replace(/\b\w/g, c => c.toUpperCase())
}

// formata CPF: 00000000000 → 000.000.000-00
function formatCPF(cpf) {
    return cpf.replace(/(\d{3})(\d{3})(\d{3})(\d{2})/, "$1.$2.$3-$4")
}

// formata telefone: 24987654321 → (24) 98765-4321
function formatTel(tel) {
    if (!tel) return "—"
    const t = tel.replace(/\D/g, "")
    if (t.length === 11) return t.replace(/(\d{2})(\d{5})(\d{4})/, "($1) $2-$3")
    if (t.length === 10) return t.replace(/(\d{2})(\d{4})(\d{4})/, "($1) $2-$3")
    return tel
}

// formata data: 2020-02-11 → 11/02/2020
function formatDate(dateStr) {
    if (!dateStr) return "—"
    const d = dateStr.split("T")[0]
    const [year, month, day] = d.split("-")
    return `${day}/${month}/${year}`
}

// logout
function logout() {
    localStorage.removeItem("token")
    localStorage.removeItem("admin")
    window.location.href = "../index.html"
}

// máscara de CPF
function maskCPF(input) {
    let v = input.value.replace(/\D/g, "")
    v = v.replace(/(\d{3})(\d)/, "$1.$2")
    v = v.replace(/(\d{3})(\d)/, "$1.$2")
    v = v.replace(/(\d{3})(\d{1,2})$/, "$1-$2")
    input.value = v
}

// máscara de telefone
function maskTel(input) {
    let v = input.value.replace(/\D/g, "")
    v = v.replace(/(\d{2})(\d)/, "($1) $2")
    v = v.replace(/(\d{5})(\d{1,4})$/, "$1-$2")
    input.value = v
}

// inicializa
loadAdminName()
loadAssociate()
// abre modal de pagamento com valores padrão
function openPaymentModal() {
    const today = new Date()
    const year = today.getFullYear()
    const month = String(today.getMonth() + 1).padStart(2, "0")
    const day = String(today.getDate()).padStart(2, "0")

    document.getElementById("inputCompetence").value = `${year}-${month}`
    document.getElementById("inputPaymentDate").value = `${year}-${month}-${day}`
    document.getElementById("inputValue").value = ""
    document.getElementById("inputPaymentStatus").value = "true"
    document.getElementById("paymentModalError").textContent = ""
    document.getElementById("paymentModalOverlay").classList.add("active")
}

// fecha modal de pagamento
function closePaymentModal() {
    document.getElementById("paymentModalOverlay").classList.remove("active")
}

// fecha modal de pagamento ao clicar fora
function closePaymentModalOnOverlay(event) {
    if (event.target === document.getElementById("paymentModalOverlay")) {
        closePaymentModal()
    }
}

// registra pagamento do associado
async function savePayment() {
    const token = checkAuth()
    const errorEl = document.getElementById("paymentModalError")
    const btnSave = document.getElementById("btnSavePayment")

    errorEl.textContent = ""

    const competenceRaw = document.getElementById("inputCompetence").value
    const paymentDate = document.getElementById("inputPaymentDate").value
    const value = parseFloat(document.getElementById("inputValue").value)
    const status = document.getElementById("inputPaymentStatus").value === "true"

    if (!competenceRaw) {
        errorEl.textContent = "Competência é obrigatória."
        return
    }

    if (!paymentDate) {
        errorEl.textContent = "Data do pagamento é obrigatória."
        return
    }

    if (!value || value <= 0) {
        errorEl.textContent = "Valor deve ser maior que zero."
        return
    }

    const body = {
        associateID: currentAssociate.id,
        competence: `${competenceRaw}-01`,
        paymentDate: paymentDate,
        value: value,
        status: status
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
            body: JSON.stringify(body)
        })

        const data = await response.json()

        if (!response.ok) {
            errorEl.textContent = data.error || "Erro ao registrar pagamento."
            return
        }

        closePaymentModal()
        alert(`Pagamento registrado com sucesso!`)

    } catch (error) {
        errorEl.textContent = "Não foi possível conectar ao servidor."
    } finally {
        btnSave.disabled = false
        btnSave.textContent = "Registrar"
    }
}

// busca o histórico de pagamentos do associado
async function loadPaymentHistory() {
    const token = checkAuth()
    const id = getIDFromURL()
    const tbody = document.getElementById("paymentHistoryTable")

    try {
        const response = await fetch(`${API_URL}/payments/associate/${id}`, {
            method: "GET",
            headers: { "Authorization": `Bearer ${token}` }
        })

        if (response.status === 401) { logout(); return }

        const payments = await response.json()
        renderPaymentHistory(payments)

    } catch (error) {
        tbody.innerHTML = `<tr><td colspan="4" class="table-loading">Erro ao carregar histórico.</td></tr>`
    }
}

// renderiza o histórico de pagamentos
function renderPaymentHistory(payments) {
    const tbody = document.getElementById("paymentHistoryTable")

    if (payments.length === 0) {
        tbody.innerHTML = `<tr><td colspan="4" class="table-loading">Nenhum pagamento registrado.</td></tr>`
        return
    }

    tbody.innerHTML = payments.map(p => `
        <tr>
            <td>${formatDate(p.competence)}</td>
            <td>${formatDate(p.paymentDate)}</td>
            <td>${formatCurrency(p.value)}</td>
            <td>
                <span class="badge ${p.status ? 'badge-paid' : 'badge-pending'}">
                    ${p.status ? 'Pago' : 'Pendente'}
                </span>
            </td>
        </tr>
    `).join("")
}

// formata valor em reais
function formatCurrency(value) {
    return value.toLocaleString("pt-BR", {
        style: "currency",
        currency: "BRL"
    })
}
loadPaymentHistory()