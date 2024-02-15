# Profitability Golang

## ctrl+s :v:

O Profitability é uma API RESTful que fornece a rentabilidade líquida e a tributação correspondente para investimentos de renda fixa, como CBD, CRI, LC, LCA e LCI, pré ou pós-fixados, com referência à SELIC ou ao IPCA.


## Funcionalidades:

- A API RESTful oferece cinco rotas principais:
  - Recuperação de dados em cache obtidos de consultas a dados diretamente obtidos do Banco Central através de sua [API](https://www.bcb.gov.br/).
  - Cálculo da rentabilidade líquida considerando o desconto de impostos, a partir da taxa bruta nominal do contrato e da duração do contrato.


## Configuração e Inicialização:

- A aplicação carrega a configuração do ambiente, incluindo a porta na qual o servidor será executado.
- Utiliza o pacote `mux` para roteamento HTTP e `cors` para lidar com solicitações CORS.
- Inicializa serviços, controladores e cache.
- Integra-se com APIs externas para obter dados, como as taxas SELIC e IPCA.


## Uso

Para usar a API de Rentabilidade, os desenvolvedores precisam enviar solicitações HTTP para os pontos finais designados. A API responde com dados JSON contendo as informações solicitadas ou valores de moeda convertidos.


## API Endpoints


- **Taxas Atuais:**

  Este endpoint retorna as taxas de juros SELIC e IPCA utilizadas nos cálculos de rentabilidade.
  
  Endpoint:
  
  GET /api/taxas
  
  Exemplo de Uso:
  
  ```bash
  curl "http://localhost:8000/api/taxas"
  ```
  
  Este exemplo retorna as taxas SELIC e IPCA em formato JSON.


Para obter informações mais detalhadas sobre as rotas e suas funcionalidades, consulte a [documentação central da API](https://rmottanet.gitbook.io/profitability).


## Contribuições

Contribuições para o projeto da API de Rentabilidade são bem-vindas! Se você tiver ideias para melhorias, solicitações de funcionalidades ou relatórios de bugs, sinta-se à vontade para abrir um problema ou enviar um pull request.

Obrigado por considerar a API de Rentabilidade para suas necessidades de cálculo. Se você tiver alguma dúvida ou precisar de mais assistência, não hesite em entrar em contato. Feliz codificação! 🚀

<br />
<br />
<p align="center">
<a href="https://gitlab.com/rmotta.net"><img src="https://img.shields.io/badge/Gitlab--_.svg?style=social&logo=gitlab" alt="GitLab"></a>
<a href="https://github.com/rmottanet"><img src="https://img.shields.io/badge/Github--_.svg?style=social&logo=github" alt="GitHub"></a>
<a href="https://instagram.com/rmottanet/"><img src="https://img.shields.io/badge/Instagram--_.svg?style=social&logo=instagram" alt="Instagram"></a>
<a href="https://www.linkedin.com/in/rmottanet/"><img src="https://img.shields.io/badge/Linkedin--_.svg?style=social&logo=linkedin" alt="Linkedin"></a>
</p>
<br />
