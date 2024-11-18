function getTopProduct() {
  return  fetch('https://fakestoreapi.com/products')
    .then((res) => res.json())
    .then((jsonData) => {
      return jsonData
    })
    .catch((err) => {
      console.log(err)
    })
}

function getCategory() {
  return  fetch('https://fakestoreapi.com/products/categories')
    .then((res) => res.json())
    .then((jsonData) => {
      return jsonData
    })
    .catch((err) => {
      console.log(err)
    })
}

function getListProductCategory(category) {
  return  fetch(`https://fakestoreapi.com/products/category/${category}`)
    .then((res) => res.json())
    .then((jsonData) => {
      return jsonData
    })
    .catch((err) => {
      console.log(err)
    })
}
export { getTopProduct ,getCategory,getListProductCategory }


