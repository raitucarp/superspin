package superspin

import "testing"

var superspin SuperSpin

func init() {
	superspin = New()
}

func TestCreateFourBlock(t *testing.T) {
	text := "Hello, {You {should} ya} banned {me}"

	superspin.Spin(text)

	if len(superspin.blocks) != 4 {
		t.Fatal("Blocks is not equal to 4")
	}
}

var str = "Hello there. {You should eat {fruits, veggies and grains|veggies, fruits and grains|fruits, grains and veggies|grains, fruits and veggies|veggies, grains and fruits|grains, veggies and fruits} if you want to be healthy.|If you want to be healthy, you should eat {fruits, veggies and grains|veggies, fruits and grains|fruits, grains and veggies|grains, fruits and veggies|veggies, grains and fruits|grains, veggies and fruits}.}"

func TestSpinText(t *testing.T) {
	spinText := superspin.Spin(str)
	if spinText == str {
		t.Fatal("Spin not successful")
	}
}

// Do multiple spin, with different results
func TestDoMultipleSpin(t *testing.T) {
	spinText := superspin.Spin(str)
	spinText2 := superspin.Spin(str)
	spinText3 := superspin.Spin(str)

	if spinText == str {
		t.Fatal("Spun text is equal original text. Thus spin fail.")
	}

	if spinText == spinText2 {
		t.Fatal("Spun text 1 is equal spin text 2. Thus spin fail.")
	}

	if spinText == spinText3 {
		t.Fatal("Spun text 1 is equal spin text 3. Thus spin fail.")
	}

	if spinText2 == spinText3 {
		t.Fatal("Spun text 2 is equal spin text 3. Thus spin fail.")
	}
}

// Create static spin text. Do multiple spin, but has same results
func TestStaticSpin(t *testing.T) {
	superspin.Seed(1234)
	spinText := superspin.Spin(str)
	spinText2 := superspin.Spin(str)

	if spinText != spinText2 {
		t.Fatal("Static seed fail")
	}
}
