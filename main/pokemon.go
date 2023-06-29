package main

//import (
//	"fmt"
//	"github.com/mtslzr/pokeapi-go"
//)
//
///*
//
//Call the pokemon endpoint, get first 100 pokemon
//
//return just the name,height,weight
//
//filter weight > 50
//
//sum the height of those animals
//print the height
// */
//
//type pokemon struct {
//	name string
//	weight int
//	height int
//}
//
//func main(){
//
//	r, _ := pokeapi.Resource("pokemon", 0, 100)
//	pokemonMap := make(map[string]pokemon, 100)
//	heights := 0
//	for _, aPokemon := range r.Results {
//		thePokemon, _ := pokeapi.Pokemon(aPokemon.Name)
//		pokemonMap[thePokemon.Name] = pokemon{name: thePokemon.Name,
//			weight: thePokemon.Weight,
//		height: thePokemon.Height}
//		if FilterWeight(thePokemon.Weight) {
//			heights += thePokemon.Height
//		}
//		// call filter weight with the weight, if its greater than 50 return true
//		// if true add it to the pokemonMap
//	}
//	fmt.Println(pokemonMap)
//	fmt.Print("The combined height is: ")
//	print(heights)
//}
//
//func FilterWeight(weight int) bool {
//	if weight > 50 {
//		return true
//	}
//	return false
//}