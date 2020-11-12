package main


var wbuilder = NewWeaponBuilder()
var Knife = wbuilder.setName("knife").setDamage(2).Create()
var Sword = wbuilder.setName("sword").setDamage(4).Create()
var Axe = wbuilder.setName("axe").setDamage(5).Create()

