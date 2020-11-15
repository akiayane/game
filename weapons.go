package main

import "fmt"

type weapon interface {
	GetDamage() int
	GetName() string
	GetPrice() int
	getInfo()
}

type Weapon struct {
	damage int
	name string
	price int
}

func (w *Weapon) GetName() string {
	return w.name
}

func (w *Weapon) GetDamage() int{
	return w.damage
}

func (w *Weapon) GetPrice() int {
	return w.price
}

func (w Weapon) getInfo()  {
	fmt.Println("Name -",w.name)
	fmt.Println("Damage -",w.damage)
	fmt.Println("Price:", w.price)
	fmt.Println()
}


type WeaponBuilder interface {
	setName(name string) WeaponBuilder
	setDamage(num int) WeaponBuilder
	setPrice(num int) WeaponBuilder
	Create() weapon
}

type builder struct {
	damage int
	name string
	price int
}

func NewWeaponBuilder() *builder {
	return &builder{}
}

func (cb *builder) setDamage(d int) WeaponBuilder {
	cb.damage = d
	return cb
}

func (cb *builder) setName(n string) WeaponBuilder {
	cb.name=n
	return cb
}

func (cb *builder) setPrice(p int) WeaponBuilder {
	cb.price = p
	return cb
}

func (cb *builder) Create() weapon {
	return &Weapon{
		cb.damage,
		cb.name,
		cb.price,
	}
}