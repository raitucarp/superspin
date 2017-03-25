# Introduction

Superspin is a JavaScript library for spinning a text. This is useful for article generator that seed many alternative texts.

# Installation

```
npm install superspin
```

# How to use
```javascript
const Superspin = require('Superspin');


const $superspin = new Superspin();

let str = 'Hello there. {You should eat {fruits, veggies and grains|veggies, fruits and grains|fruits, grains and veggies|grains, fruits and veggies|veggies, grains and fruits|grains, veggies and fruits} if you want to be healthy.|If you want to be healthy, you should eat {fruits, veggies and grains|veggies, fruits and grains|fruits, grains and veggies|grains, fruits and veggies|veggies, grains and fruits|grains, veggies and fruits}.}';

let spinText = $superspin.spin(str);
console.log(spinText)
```
That's it

# Methods
## spin(String text[, Boolean withSeed])

Do spin a text with random seed. If withSeed set to true, then it will generate static result based on seed. Useful for a unique spun text.

## seed(String str)

Set superspin seed. This seed will be used as seed generator if spin second argument set to true.

# License

MIT
