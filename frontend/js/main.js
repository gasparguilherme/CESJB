const API_URL = "http://localhost:8088"

async function login() {
    const email = document.getElementById("email").value
    const password = document.getElementById("password").value
    const errorMessage = document.getElementById("errorMessage")
    const btnLogin = document.getElementById("btnLogin")

    // limpa erro anterior
    errorMessage.textContent = ""

    // validação básica antes de chamar a API
    if (!email || !password) {
        errorMessage.textContent = "Preencha e-mail e senha."
        return
    }

    // desabilita o botão enquanto aguarda a resposta da API
    btnLogin.disabled = true
    btnLogin.textContent = "Entrando..."

    try {
        const response = await fetch(`${API_URL}/login`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({ email, password })
        })

        const data = await response.json()

        if (!response.ok) {
            // API retornou erro (401, 400, etc)
            errorMessage.textContent = data.error || "Erro ao realizar login."
            return
        }

        // login ok — salva o token no localStorage
        localStorage.setItem("token", data.token)
        localStorage.setItem("admin", JSON.stringify(data.admin))

        // redireciona para o dashboard
        window.location.href = "pages/dashboard.html"

    } catch (error) {
        // erro de conexão — API fora do ar ou CORS
        errorMessage.textContent = "Não foi possível conectar ao servidor."
    } finally {
        // reabilita o botão independente do resultado
        btnLogin.disabled = false
        btnLogin.textContent = "Entrar"
    }
}

// permite fazer login pressionando Enter
document.addEventListener("keydown", function (e) {
    if (e.key === "Enter") {
        login()
    }
})