function getTopProduct() {
  return  fetch('http://luquin.test:8080/api/products/list-products')
    .then((res) => res.json())
    .then((jsonData) => {
      return jsonData
    })
    .catch((err) => {
      console.log(err)
    })
}

function getCategory() {
  return  fetch('http://luquin.test:8080/api/products/categories')
    .then((res) => res.json())
    .then((jsonData) => {
      return jsonData
    })
    .catch((err) => {
      console.log(err)
    })
}

function getListProductCategory(category) {
  return  fetch(`http://luquin.test:8080/api/products/category/${category}`)
    .then((res) => res.json())
    .then((jsonData) => {
      return jsonData
    })
    .catch((err) => {
      console.log(err)
    })
}
export { getTopProduct ,getCategory,getListProductCategory }


