# 🌤️ Go Weather API - Consulta de Temperatura por CEP

Esta é uma API escrita em Go que recebe um **CEP brasileiro** como parâmetro, utiliza o serviço da **Geoapify** para converter o CEP em latitude e longitude, e em seguida consulta a **WeatherAPI** para retornar a **temperatura atual** da localidade.

## 🔗 Link de demonstração (Google Cloud Run)

> ⚠️ **Ambiente temporário.** Esta API está publicada no Google Cloud Run para fins de avaliação.

✅ Acesse a API aqui:  
- [https://go-weather-779180261777.us-east1.run.app/weather?cep=88010-040](https://go-weather-779180261777.us-east1.run.app/weather?cep=88010-040)

---

## 🚀 Tecnologias Utilizadas

| Categoria       | Tecnologia                       |
|----------------|-----------------------------------|
| Linguagem       | [Go (Golang)](https://golang.org/) |
| Framework Web   | [net/http](https://pkg.go.dev/net/http) (padrão do Go) |
| Geolocalização  | [Geoapify API](https://www.geoapify.com/) |
| Clima           | [WeatherAPI](https://www.weatherapi.com/) |
| Infraestrutura  | [Google Cloud Run](https://cloud.google.com/run) |
| Containerização | [Docker](https://www.docker.com/) + [Docker Compose](https://docs.docker.com/compose/) |
| CI/CD           | [GitHub Actions](https://github.com/features/actions) |

---

## 📦 Como Executar com Docker Compose

1. **Clone o repositório:**

```bash
git clone https://github.com/ericoalmeida/go-wheather.git
cd go-weather-api
```

2. **Crie um arquivo .env na raiz com as seguintes variáveis:**

```bash
GEOAPIFY_BASE_URL=https://api.geoapify.com
GEOAPIFY_API_KEY=<SetYourKey>
WEATHER_BASE_URL=http://api.weatherapi.com
WEATHER_API_KEY=<SetYourKey>
```

3. **Execute a aplicação com Docker Compose:**

```bash
docker-compose up --build
```

4. **Acesse a API localmente em:**

```bash
http://localhost:8080/weather?cep=78590-000
```

## Testes

1. Execute o comando abaixo para rodar os testes

```bash
go test ./... -v
```
