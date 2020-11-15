package main

var mainName string

var wbuilder = NewWeaponBuilder()
var Knife = wbuilder.setName("knife").setDamage(5).setPrice(80).Create()
var Sword = wbuilder.setName("sword").setDamage(8).setPrice(120).Create()
var Axe = wbuilder.setName("axe").setDamage(15).setPrice(200).Create()
var Excalibur = wbuilder.setName("excalibur").setDamage(200).setPrice(10000).Create()

var weather = []string{"sun", "normal", "rain", "snow"}

var concreteWeather weatherCondition