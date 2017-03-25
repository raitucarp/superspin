import test from 'ava';
import Superspin from './index.js';

const $superspin = new Superspin();

test('Create four block', t => {
  let str = 'Hello, {You {should} ya} banned {me}';

  $superspin.spin(str);
  t.is($superspin.blocksLength(), 4);
});

let str = 'Hello there. {You should eat {fruits, veggies and grains|veggies, fruits and grains|fruits, grains and veggies|grains, fruits and veggies|veggies, grains and fruits|grains, veggies and fruits} if you want to be healthy.|If you want to be healthy, you should eat {fruits, veggies and grains|veggies, fruits and grains|fruits, grains and veggies|grains, fruits and veggies|veggies, grains and fruits|grains, veggies and fruits}.}';

test('Spin a text', t => {
  let spinText = $superspin.spin(str);
  t.not(spinText, str);
});


test('Do multiple spin, with different results', t => {
  let spinText = $superspin.spin(str);
  let spinText2 = $superspin.spin(str);
  let spinText3 = $superspin.spin(str);
  t.not(spinText, str);
  t.not(spinText, spinText2);
  t.not(spinText, spinText3);
  t.not(spinText2, spinText3);
});


test('Create static spin text. Do multiple spin, but has same results', t => {
  $superspin.seed('What the fuck');

  let spinText = $superspin.spin(str, true);
  let spinText2 = $superspin.spin(str, true);

  t.is(spinText, spinText2);
});

// test travis here
