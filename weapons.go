package main

import "fmt"

type weapon interface {
	GetDamage() int
	getInfo()
}

type Weapon struct {
	damage int
	name string
}


func (cb *builder) setDamage(d int) WeaponBuilder {
	cb.damage = d
	return cb
}

func (w *Weapon) GetName() string {
	return w.name
}

func (w *Weapon) GetDamage()  int{
	return w.damage
}

func (cb *builder) setName(n string) WeaponBuilder {
	cb.name=n
	return cb
}

func (w Weapon) getInfo()  {
	fmt.Println("Damage -",w.damage)
	fmt.Println("Name -",w.damage)
}

type WeaponBuilder interface {
	setName(name string) WeaponBuilder
	setDamage(num int) WeaponBuilder
	Create() weapon
}

type builder struct {
	damage int
	name string
}

func NewWeaponBuilder() *builder {
	return &builder{}
}

func (cb *builder) Create() weapon {
	return &Weapon{
		cb.damage,
		cb.name,
	}
}