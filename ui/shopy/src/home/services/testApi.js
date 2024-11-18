function getTopProduct() {
  return  fetch('https://fakestoreapi.com/products?limit=6')
    .then((res) => res.json())
    .then((jsonData) => {
      return jsonData
    })
    .catch((err) => {
      console.log(err)
    })
}

export { getTopProduct }
