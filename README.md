🚀 Aplicação Web em Golang
Um projeto de aplicação web desenvolvido em Go para gerenciamento de produtos, seguindo as melhores práticas de organização de código.

✨ Funcionalidades
Este projeto implementa um CRUD completo para a entidade "Produtos". Com ele, você pode:

✅ Criar novos produtos.

✏️ Editar informações de produtos existentes.

🗑️ Deletar produtos do banco de dados.

🔄 Atualizar o status ou qualquer outro dado de um produto.

📄 Ler e listar todos os produtos cadastrados.

🏛️ Arquitetura
O projeto foi estruturado para ser modular, escalável e de fácil manutenção, utilizando o padrão MVC (Model-View-Controller).

/meu-projeto
├── controllers/
│   └── product_controller.go  # Lógica de manipulação das requisições
├── models/
│   └── product.go             # Estrutura de dados e acesso ao banco
├── views/
│   └── (html, templates, etc) # Camada de apresentação
├── routes/
│   └── router.go              # Módulo dedicado para o gerenciamento de todas as rotas da API
└── main.go                    # Ponto de entrada da aplicação
📦 Módulos Organizados: A aplicação é dividida em pacotes claros e com responsabilidades bem definidas.

🛣️ Rotas Dedicadas: Um módulo exclusivo (/routes) para centralizar e gerenciar todos os endpoints da aplicação, tornando o código mais limpo e fácil de encontrar.

🧠 Controllers Inteligentes: Os controllers (/controllers) são responsáveis por intermediar as requisições, aplicando a lógica de negócio necessária antes de interagir com o model ou a view.

🚀 Como Executar (Exemplo)
Para rodar este projeto, você precisará ter o Go instalado.
# Clone o repositório
git clone https://github.com/EduardoBotelho/Loja_Golang.git

# Navegue até o diretório do projeto
cd EduardoBotelho/Loja_Golang

# Instale as dependências (se houver)
go mod tidy

# Execute a aplicação
go run main.go
