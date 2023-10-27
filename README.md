<p align="center">
  <a href="https://www.truemoney.com/" target="blank"><img src="https://www.truemoney.com/wp-content/uploads/2022/01/truemoneywallet-sendgift-hongbao-20220125-icon-2.png" width="200" height="200" alt="Angpao Logo" /></a>
</p>

## 👋 Description
This project allows you to easily and quickly host an API for redeeming TrueWallet vouchers. You only need to enter the voucher code and mobile number.

## 🧃 Preface

<p>I want to create an API that can easily redeem TrueWallet vouchers and ensure fast processing speeds. To achieve this, I am using Golang to develop the project.</p>

## 📝 How to use?

- Clone this project following the installation instructions.
- Setting the port in the .env file.
- Installing all the necessary packages.
- Run the program.

## 📚 Installation

```bash
# Clone project
$ git clone https://github.com/jumpogpo/golang-truewallet-api.git
$ cd golang-truewallet-api

# Install packages
$ go mod download
```

## 📺 Running the app

```bash
# run
$ go run .

# build
$ go build .
```

## ▶️ How to use
- Method: `Post`
- Url: `http://localhost:{port}/redeem/{angpao_code}`
- Body:
```json
{
  "mobile": "phoneNumber"
}
```

## 🤝 Reference

- TrueWallet - [https://www.truemoney.com/](https://www.truemoney.com/)
- Golang - [https://go.dev/](https://go.dev/)
- Fiber - [https://docs.gofiber.io/](https://docs.gofiber.io/)
