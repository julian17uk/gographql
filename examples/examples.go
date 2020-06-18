// To test, run GraphiQL and visit endpoint http://localhost:8080/graphql?
// with the following graphql query
// query  {
//	readFile(
//		filename: "hello"
//	)
// }
// To test graphql mutation
// mutation {
// 	createNewFile(
// 	  filename: "australia",
// 	  text: "hello from australia",
// 	)
//   }
//
// mutation {
// 	addToFile(
// 	  filename: "australia",
// 	  text: "major city is Melbourne",
// 	)
//   }

// or using a browser with the following http request
// query
// http://localhost:8080/graphql?query={readFile(filename:"hello")}
// or
// mutations
// http://[ipv6address]:8080/graphql?query=mutation{createNewFile(filename:"france", text:"major city is Sao Paulo")}
// http://http://localhost:8080/graphql?query=mutation{addToFile(filename:%22france%22,text:%22a%20major%20city%20is%20bordeaux%22)}
