const API_URL = "http://localhost:8088"

let allAssociates = [] // guarda todos os associados para busca local

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

// busca todos os associados e preenche a tabela
async function loadAssociates() {
    const token = checkAuth()

    try {
        const response = await fetch(`${API_URL}/associates`, {
            method: "GET",
            headers: { "Authorization": `Bearer ${token}` }
        })

        if (response.status === 401) { logout(); return }

        allAssociates = await response.json()
        renderTable(allAssociates)

    } catch (error) {
        document.getElementById("associatesTable").innerHTML = `
            <tr><td colspan="6" class="table-loading">Erro ao carregar associados.</td></tr>
        `
    }
}

// renderiza a tabela com a lista recebida
function renderTable(associates) {
    const tbody = document.getElementById("associatesTable")

    if (associates.length === 0) {
        tbody.innerHTML = `
            <tr><td colspan="6" class="table-loading">Nenhum associado encontrado.</td></tr>
        `
        return
    }

    tbody.innerHTML = associates.map(a => `
        <tr>
            <td><a href="associate_detail.html?id=${a.id}" class="associate-link">${a.name}</a></td>
            <td>${formatCPF(a.cpf)}</td>
            <td>${a.email}</td>
            <td>${a.position}</td>
            <td>
                <span class="badge ${a.status ? 'badge-active' : 'badge-inactive'}">
                    ${a.status ? 'Ativo' : 'Inativo'}
                </span>
            </td>
            <td>
                <button class="btn-edit" onclick="openModalEdit(${a.id})" title="Editar">✏️</button>
                <button class="btn-deactivate" onclick="deactivateAssociate(${a.id})" title="Desativar">🚫</button>
            </td>
        </tr>
    `).join("")
}

// busca local por nome ou CPF
function handleSearch() {
    const query = document.getElementById("searchInput").value.toLowerCase().trim()

    if (query === "") {
        renderTable(allAssociates)
        return
    }

    const filtered = allAssociates.filter(a =>
        a.name.toLowerCase().includes(query) ||
        a.cpf.includes(query.replace(/\D/g, ""))
    )

    renderTable(filtered)
}

// abre modal para novo associado
function openModal() {
    document.getElementById("modalTitle").textContent = "Novo Associado"
    document.getElementById("associateId").value = ""
    document.getElementById("inputName").value = ""
    document.getElementById("inputCPF").value = ""
    document.getElementById("inputEmail").value = ""
    document.getElementById("inputTel").value = ""
    document.getElementById("inputDateOfBirth").value = ""
    document.getElementById("inputAssociationDate").value = ""
    document.getElementById("inputPosition").value = ""
    document.getElementById("inputAddress").value = ""
    document.getElementById("inputStatus").value = "true"
    document.getElementById("modalError").textContent = ""
    document.getElementById("modalOverlay").classList.add("active")
}

// abre modal para editar associado existente
function openModalEdit(id) {
    const associate = allAssociates.find(a => a.id === id)
    if (!associate) return

    document.getElementById("modalTitle").textContent = "Editar Associado"
    document.getElementById("associateId").value = associate.id
    document.getElementById("inputName").value = associate.name
    document.getElementById("inputCPF").value = formatCPF(associate.cpf)
    document.getElementById("inputEmail").value = associate.email
    document.getElementById("inputTel").value = associate.tel
    document.getElementById("inputDateOfBirth").value = associate.date_of_birth?.split("T")[0] || ""
    document.getElementById("inputAssociationDate").value = associate.association_date?.split("T")[0] || ""
    document.getElementById("inputPosition").value = associate.position
    document.getElementById("inputAddress").value = associate.address
    document.getElementById("inputStatus").value = associate.status ? "true" : "false"
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

// salva associado (cria ou edita)
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

    btnSave.disabled = true
    btnSave.textContent = "Salvando..."

    try {
        const isEdit = id !== ""
        const url = isEdit ? `${API_URL}/associate/${id}` : `${API_URL}/associate`
        const method = isEdit ? "PUT" : "POST"

        const response = await fetch(url, {
            method,
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
        loadAssociates()

    } catch (error) {
        errorEl.textContent = "Não foi possível conectar ao servidor."
    } finally {
        btnSave.disabled = false
        btnSave.textContent = "Salvar"
    }
}

// desativa associado
async function deactivateAssociate(id) {
    if (!confirm("Deseja desativar este associado?")) return

    const token = checkAuth()

    try {
        const associate = allAssociates.find(a => a.id === id)
        if (!associate) return

        const body = {
            ...associate,
            cpf: associate.cpf.replace(/\D/g, ""),
            date_of_birth: associate.date_of_birth?.split("T")[0],
            association_date: associate.association_date?.split("T")[0],
            status: false
        }

        const response = await fetch(`${API_URL}/associate/${id}`, {
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

        loadAssociates()

    } catch (error) {
        alert("Não foi possível conectar ao servidor.")
    }
}

// converte string para Title Case: "joao silva" → "Joao Silva"
function toTitleCase(str) {
    return str.toLowerCase().replace(/\b\w/g, c => c.toUpperCase())
}

// máscara de CPF: 000.000.000-00
function maskCPF(input) {
    let v = input.value.replace(/\D/g, "")
    v = v.replace(/(\d{3})(\d)/, "$1.$2")
    v = v.replace(/(\d{3})(\d)/, "$1.$2")
    v = v.replace(/(\d{3})(\d{1,2})$/, "$1-$2")
    input.value = v
}

// máscara de telefone: (00) 00000-0000
function maskTel(input) {
    let v = input.value.replace(/\D/g, "")
    v = v.replace(/(\d{2})(\d)/, "($1) $2")
    v = v.replace(/(\d{5})(\d{1,4})$/, "$1-$2")
    input.value = v
}

// formata CPF para exibição
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
loadAssociates()