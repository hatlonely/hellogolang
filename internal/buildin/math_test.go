package buildin

import (
	. "github.com/smartystreets/goconvey/convey"
	"math"
	"math/big"
	"testing"
)

func TestMath(t *testing.T) {
	Convey("test math", t, func() {
		So(math.Abs(-6), ShouldAlmostEqual, 6)
		So(math.Floor(-6.2), ShouldAlmostEqual, -7) // 向下取整
		So(math.Floor(6.2), ShouldAlmostEqual, 6)
		So(math.Floor(6.8), ShouldAlmostEqual, 6)
		So(math.Ceil(-6.2), ShouldAlmostEqual, -6) // 向上取整
		So(math.Ceil(6.2), ShouldAlmostEqual, 7)
		So(math.Ceil(6.8), ShouldAlmostEqual, 7)
		So(math.Round(-6.2), ShouldAlmostEqual, -6) // 四舍五入
		So(math.Round(6.2), ShouldAlmostEqual, 6)
		So(math.Round(6.8), ShouldAlmostEqual, 7)
		So(math.Round(-6.5), ShouldAlmostEqual, -7)
		So(math.Round(6.5), ShouldAlmostEqual, 7)
		So(math.RoundToEven(6.8), ShouldAlmostEqual, 7) // 返回最接近的整数，如果有两个，返回偶数
		So(math.RoundToEven(6.5), ShouldAlmostEqual, 6)
		So(math.RoundToEven(7.5), ShouldAlmostEqual, 8)
		So(math.RoundToEven(-6.5), ShouldAlmostEqual, -6)
		So(math.RoundToEven(-7.5), ShouldAlmostEqual, -8)

		So(math.Max(123, 456), ShouldAlmostEqual, 456)
		So(math.Min(123, 456), ShouldAlmostEqual, 123)
		So(math.Mod(7, 3), ShouldAlmostEqual, 1) // floor 取模
		So(math.Mod(-7, 3), ShouldAlmostEqual, -1)
		So(math.Mod(7, -3), ShouldAlmostEqual, 1)
		So(math.Mod(8, 3), ShouldAlmostEqual, 8-math.Floor(8.0/3.0)*3.0)
		So(math.Remainder(7, 3), ShouldAlmostEqual, 1) // round 取模
		So(math.Remainder(-7, 3), ShouldAlmostEqual, -1)
		So(math.Remainder(7, -3), ShouldAlmostEqual, 1)
		So(math.Remainder(8, 3), ShouldAlmostEqual, 8-math.Round(8.0/3.0)*3.0)
		i, f := math.Modf(123.456)
		So(i, ShouldAlmostEqual, 123)
		So(f, ShouldAlmostEqual, 0.456)

		So(math.Sin(math.Pi/2), ShouldAlmostEqual, 1)
		So(math.Cos(math.Pi), ShouldAlmostEqual, -1)
		So(math.Tan(math.Pi/4), ShouldAlmostEqual, 1)
		So(math.Asin(1), ShouldAlmostEqual, math.Pi/2)
		So(math.Acos(-1), ShouldAlmostEqual, math.Pi)
		So(math.Atan(1), ShouldAlmostEqual, math.Pi/4)
		So(math.Atan2(1.0, 4.0), ShouldAlmostEqual, math.Atan(0.25))
		sin, cos := math.Sincos(math.Pi / 4)
		So(sin, ShouldAlmostEqual, math.Sin(math.Pi/4))
		So(cos, ShouldAlmostEqual, math.Cos(math.Pi/4))

		So(math.Pow(3, 2), ShouldAlmostEqual, 9)
		So(math.Pow(2, 3), ShouldAlmostEqual, 8)
		So(math.Pow10(4), ShouldAlmostEqual, 10000)
		So(math.Sqrt(4), ShouldAlmostEqual, 2)
		So(math.Cbrt(27), ShouldAlmostEqual, 3)
		So(math.Hypot(3, 4), ShouldAlmostEqual, 5) // √(x² + y²)

		So(math.Exp(2.5), ShouldAlmostEqual, math.Pow(math.E, 2.5)) // e ^ x
		So(math.Expm1(2), ShouldAlmostEqual, math.Exp(2)-1)         // e ^ x - 1
		So(math.Log(math.Exp(1.5)), ShouldAlmostEqual, 1.5)         // ln(x)
		So(math.Log10(1000), ShouldAlmostEqual, 3)                  // lg(x)
		So(math.Log1p(math.E-1), ShouldAlmostEqual, 1)              // ln(1+x)
		So(math.Log2(1024), ShouldAlmostEqual, 10)
		So(math.Logb(1234), ShouldAlmostEqual, math.Floor(math.Log2(1234)))
		So(math.Ilogb(1234), ShouldEqual, int(math.Floor(math.Log2(1234))))
		f, e := math.Frexp(1024) // 将一个数写成 f * 2 ^ e 的形式
		So(f, ShouldAlmostEqual, 0.5)
		So(e, ShouldAlmostEqual, 11)
		So(math.Ldexp(f, e), ShouldAlmostEqual, 1024)

		So(math.Sinh(2), ShouldAlmostEqual, (math.Exp(2)-math.Exp(-2))/2) // sinh(x) = (e ^ x - e ^ -x) / 2
		So(math.Cosh(2), ShouldAlmostEqual, (math.Exp(2)+math.Exp(-2))/2) // cosh(x) = (e ^ x + e ^ -x) / 2
		So(math.Tanh(2), ShouldAlmostEqual, math.Sinh(2)/math.Cosh(2))    // tanh(x) = sinh(x) / cosh(x)

		So(math.Dim(1, 2), ShouldAlmostEqual, math.Max(1-2, 0))
		So(math.Dim(2, 1), ShouldAlmostEqual, math.Max(2-1, 0))
		So(math.Erf(0.4), ShouldAlmostEqual, 0.42839235504666845)    // 误差函数
		So(math.Erfc(0.4), ShouldAlmostEqual, 0.5716076449533315)    // 余补误差函数
		So(math.Erfinv(0.4), ShouldAlmostEqual, 0.370807158593558)   // 逆误差函数
		So(math.Erfcinv(0.4), ShouldAlmostEqual, 0.5951160814499948) // 逆余补误差函数
		So(math.Gamma(6), ShouldAlmostEqual, 5*4*3*2*1)              // 伽马函数
		l, s := math.Lgamma(6)
		So(l, ShouldAlmostEqual, math.Log(math.Gamma(6)))
		So(s, ShouldEqual, 1)
		So(math.J0(0.5), ShouldEqual, 0.9384698072408129)     // 0阶贝塞尔函数
		So(math.J1(0.5), ShouldEqual, 0.2422684576748739)     // 1阶贝塞尔函数
		So(math.Jn(2, 0.5), ShouldEqual, 0.03060402345868264) // n阶贝塞尔函数

		// float 的二进制表示用整型读取
		So(math.Float32bits(123.456), ShouldEqual, 1123477881)
		So(math.Float32frombits(1123477881), ShouldAlmostEqual, 123.456, 0.00001)
		So(math.Float64bits(123.456), ShouldEqual, 4638387860618067575)
		So(math.Float64frombits(4638387860618067575), ShouldAlmostEqual, 123.456)
		So(math.Float32bits(math.Nextafter32(123.456, math.MaxFloat32)), ShouldEqual, 1123477881+1)
		So(math.Float64bits(math.Nextafter(123.456, math.MaxFloat64)), ShouldEqual, 4638387860618067575+1)
		So(math.Signbit(-1), ShouldEqual, -1 < 0)

		So(math.IsInf(math.Inf(1), 1), ShouldBeTrue) // 无穷
		So(math.IsNaN(math.NaN()), ShouldBeTrue)     // 不是一个数
	})
}

func TestBig(t *testing.T) {
	Convey("test big int", t, func() {
		i := big.NewInt(0)
		So(i.Abs(big.NewInt(-123)), ShouldResemble, big.NewInt(123))
		So(i.Neg(big.NewInt(-123)), ShouldResemble, big.NewInt(123))
		So(i.Neg(big.NewInt(123)), ShouldResemble, big.NewInt(-123))
		So(i.Add(big.NewInt(123), big.NewInt(456)), ShouldResemble, big.NewInt(123+456))
		So(i.Mul(big.NewInt(123), big.NewInt(456)), ShouldResemble, big.NewInt(123*456))
		So(i.Sub(big.NewInt(123), big.NewInt(456)), ShouldResemble, big.NewInt(123-456))
		So(i.Div(big.NewInt(456), big.NewInt(123)), ShouldResemble, big.NewInt(456/123))
		So(i.Mod(big.NewInt(123), big.NewInt(456)), ShouldResemble, big.NewInt(123%456))
		So(i.Rem(big.NewInt(123), big.NewInt(456)), ShouldResemble, big.NewInt(123%456))
		So(i.Sqrt(big.NewInt(123)), ShouldResemble, big.NewInt(int64(math.Sqrt(123))))
		So(i.Exp(big.NewInt(2), big.NewInt(4), big.NewInt(10000)), ShouldResemble, big.NewInt(int64(math.Pow(2, 4))%10000))
		So(i.GCD(nil, nil, big.NewInt(24), big.NewInt(36)), ShouldResemble, big.NewInt(12))

		So(i.And(big.NewInt(123), big.NewInt(456)), ShouldResemble, big.NewInt(123&456))
		So(i.Or(big.NewInt(123), big.NewInt(456)), ShouldResemble, big.NewInt(123|456))
		So(i.AndNot(big.NewInt(123), big.NewInt(456)), ShouldResemble, big.NewInt(123&(^456)))
		So(i.Xor(big.NewInt(123), big.NewInt(456)), ShouldResemble, big.NewInt(123^456))
		So(big.NewInt(123).Cmp(big.NewInt(456)), ShouldEqual, -1)
		So(big.NewInt(123).CmpAbs(big.NewInt(-456)), ShouldEqual, -1)
		So(big.NewInt(67).ProbablyPrime(2), ShouldBeTrue) // 是质数的概率 1 - 1/4^n
		So(i.Lsh(big.NewInt(123), 3), ShouldResemble, big.NewInt(123<<3))
		So(i.Rsh(big.NewInt(123), 3), ShouldResemble, big.NewInt(123>>3))

		// bitset
		i = big.NewInt(0)
		i.SetBit(i, 3, 1)
		So(i.Bit(3), ShouldEqual, 1)
		i.SetBit(i, 3, 0)
		So(i.Bit(3), ShouldEqual, 0)

		i.SetString("12345678", 10)
		So(i, ShouldResemble, big.NewInt(12345678))
		So(i.Text(10), ShouldEqual, "12345678")
	})
}
