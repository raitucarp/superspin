# Introduction

Superspin is a JavaScript library for spinning a text. This is useful for article generator that seed many alternative texts.

# Installation

Node:

```
npm install superspin
```

Go:
```
go get -u github.com/raitucarp/superspin
```

# How to use

Node/Javascript:
```javascript
const Superspin = require('superspin');


const $superspin = new Superspin();

let str = 'Hello there. {You should eat {fruits, veggies and grains|veggies, fruits and grains|fruits, grains and veggies|grains, fruits and veggies|veggies, grains and fruits|grains, veggies and fruits} if you want to be healthy.|If you want to be healthy, you should eat {fruits, veggies and grains|veggies, fruits and grains|fruits, grains and veggies|grains, fruits and veggies|veggies, grains and fruits|grains, veggies and fruits}.}';

let spinText = $superspin.spin(str);
console.log(spinText)
```

Golang:
```go
import "github.com/raitucarp/superspin"

func main() {
  sp := superspin.New()
  str := "Hello there. {You should eat {fruits, veggies and grains|veggies, fruits and grains|fruits, grains and veggies|grains, fruits and veggies|veggies, grains and fruits|grains, veggies and fruits} if you want to be healthy.|If you want to be healthy, you should eat {fruits, veggies and grains|veggies, fruits and grains|fruits, grains and veggies|grains, fruits and veggies|veggies, grains and fruits|grains, veggies and fruits}.}"

  spinText := superspin.Spin(str)
}
```
That's it

# Methods

## Javascript/Node.js

### constructor
```
new Superspin(openToken, closeToken, orToken)
```
default params:
- openToken = {
- closeToken = {
- orToken = {

### spin(String text[, Boolean withSeed])

Do spin a text with random seed. If withSeed set to true, then it will generate static result based on seed. Useful for a unique spun text.

### seed(String str)
Set superspin seed. This seed will be used as seed generator if spin second argument set to true.

### Golang

See godocs or superspin_test.go source

# License

MIT
