/*

The program solves a linear kinematics problem.

Formula: s = 1/2 * a * (t*t) + vo * t + so

s: displacement
a: acceleration
t: time
so: initial displacement
vo: initial velocity

Example:

a = 10
vo = 2
so = 1

time = 3 => fn(3) = 52
time = 5 =>	fn(5) = 136

*/
package kinematics

func GenDisplaceFn(a float64, vo float64, so float64) func(float64) float64 {
	fn := func(time float64) float64 {
		return (0.5 * (a * (time * time))) + (vo * time) + so
	}
	return fn
}
