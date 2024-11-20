function getListProduct() {
  return fetch('http://luquin.test:8080/api/products/list-products')
    .then((res) => res.json())
    .then((jsonData) => {
      return jsonData
    })
    .catch((err) => {
      console.log(err)
    })
}

function getCategory() {
  return fetch('http://luquin.test:8080/api/products/categories')
    .then((res) => res.json())
    .then((jsonData) => {
      return jsonData
    })
    .catch((err) => {
      console.log(err)
    })
}


function postDate(formData) {
  try {
  return   fetch(`http://luquin.test:8080/api/new-products`, {
      method: 'POST',
      body: formData,
    }).then(res => res.json())
      .then((jsonData) =>{
       return jsonData
      })
      .catch(err => console.log(err))

  } catch (error) {
    console.error('Error al subir archivos:', error)
  }
}

export { getListProduct, getCategory, postDate }
