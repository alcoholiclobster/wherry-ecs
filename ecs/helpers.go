package ecs

func Get1[T Component](entity Entity, mask ComponentMask) T {
	c := (*entity.Get(mask)).(T)

	return c
}

func Get2[T1 Component, T2 Component](entity Entity, m1 ComponentMask, m2 ComponentMask) (c1 T1, c2 T2) {
	c1 = (*entity.Get(m1)).(T1)
	c2 = (*entity.Get(m2)).(T2)

	return c1, c2
}

func Get3[T1 Component, T2 Component, T3 Component](entity Entity, m1 ComponentMask, m2 ComponentMask, m3 ComponentMask) (c1 T1, c2 T2, c3 T3) {
	c1 = (*entity.Get(m1)).(T1)
	c2 = (*entity.Get(m2)).(T2)
	c3 = (*entity.Get(m3)).(T3)

	return c1, c2, c3
}

func Get4[T1 Component, T2 Component, T3 Component, T4 Component](entity Entity, m1 ComponentMask, m2 ComponentMask, m3 ComponentMask, m4 ComponentMask) (c1 T1, c2 T2, c3 T3, c4 T4) {
	c1 = (*entity.Get(m1)).(T1)
	c2 = (*entity.Get(m2)).(T2)
	c3 = (*entity.Get(m3)).(T3)
	c4 = (*entity.Get(m4)).(T4)

	return c1, c2, c3, c4
}

func Get5[T1 Component, T2 Component, T3 Component, T4 Component, T5 Component](entity Entity, m1 ComponentMask, m2 ComponentMask, m3 ComponentMask, m4 ComponentMask, m5 ComponentMask) (c1 T1, c2 T2, c3 T3, c4 T4, c5 T5) {
	c1 = (*entity.Get(m1)).(T1)
	c2 = (*entity.Get(m2)).(T2)
	c3 = (*entity.Get(m3)).(T3)
	c4 = (*entity.Get(m4)).(T4)
	c5 = (*entity.Get(m5)).(T5)

	return c1, c2, c3, c4, c5
}
