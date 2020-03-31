package gota

import "testing"

func TestKER(t *testing.T) {
	list := []float64{20, 21, 22, 23, 22, 21}

	expList := []float64{1, 1.0 / 3, 1.0 / 3}

	ker := NewKER(3)
	var actList []float64
	for _, v := range list {
		if vOut := ker.Add(v); ker.Warmed() {
			actList = append(actList, vOut)
		}
	}

	if diff := diffFloats(expList, actList, 0.0000001); diff != "" {
		t.Errorf("unexpected floats:\n%s", diff)
	}
}

func TestKAMA(t *testing.T) {
	list := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

	// expList is generated by the following code:
	// expList, _ := talib.Cmo(list, 10, nil)
	expList := []float64{10.444444444444445, 11.135802469135802, 11.964334705075446, 12.869074836153025, 13.81615268675168, 13.871008014588556, 13.71308456353558, 13.553331356741122, 13.46599437575161, 13.4515677602438, 13.29930139347417, 12.805116570729284, 11.752584300922967, 10.036160535131103, 7.797866963961725, 6.109926091089847, 4.727736717272138, 3.5154092873734104, 2.3974496040963396}

	kama := NewKAMA(10)
	var actList []float64
	for _, v := range list {
		if vOut := kama.Add(v); kama.Warmed() {
			actList = append(actList, vOut)
		}
	}

	if diff := diffFloats(expList, actList, 0.0000001); diff != "" {
		t.Errorf("unexpected floats:\n%s", diff)
	}
}

func TestKAMAWarmCount(t *testing.T) {
	period := 9
	kama := NewKAMA(period)

	var i int
	for i = 0; i < period*10; i++ {
		kama.Add(float64(i))
		if kama.Warmed() {
			break
		}
	}

	if got, want := i, kama.WarmCount(); got != want {
		t.Errorf("unexpected warm count: got=%d want=%d", got, want)
	}
}

var BenchmarkKAMAVal float64

func BenchmarkKAMA(b *testing.B) {
	list := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	for n := 0; n < b.N; n++ {
		kama := NewKAMA(5)
		for _, v := range list {
			BenchmarkKAMAVal = kama.Add(v)
		}
	}
}