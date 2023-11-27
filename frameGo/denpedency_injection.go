package framego

// import "reflect"

// type Container struct {
// 	dependencies map[reflect.Type]interface{}
// }

// func (c *Container) Provide(dep interface{}) {
// 	depType := reflect.TypeOf(dep)
// 	c.dependencies[depType] = dep
// }

 
// func (c *Container) Invoke(fn interface{}) {
// 	fnType := reflect.TypeOf(fn)
// 	fnValue := reflect.ValueOf(fn)
 
// 	// Résoudre les dépendances et appeler la fonction
// 	c.resolveDependencies(fnType, fnValue).Call(nil)
// }

//  func (c *Container) resolveDependencies(fnType reflect.Type, fnValue reflect.Value) reflect.Value {
// 	// ... logique de résolution des dépendances ...
//  }
 
 