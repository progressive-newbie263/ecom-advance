import { cart, addToCart, calculateCartQuantity } from "../data/cart.js";

async function SearchProducts(searchTerm) {
  try {
    // Fetch products from the server with the search query
    const response = await fetch(`/search?query=${encodeURIComponent(searchTerm)}`);
    const products = await response.json();
    renderProductsGrid(products); // Render the filtered products
  } catch (error) {
    console.error('Error fetching products:', error);
  }
}

function updateCartQuantity() {
  const cartQuantity = calculateCartQuantity();
  document.querySelector('.js-cart-quantity').innerHTML = cartQuantity;
}

// Function to show the "Added to Cart" message
function showingMessage(productId) {
  const addedMessage = document.querySelector(`.js-added-to-cart-${productId}`);
  addedMessage.classList.add('unhidden');
  
  // Clear previous timeout if it exists
  const previousTimeoutId = addedMessageTimeouts[productId];
  if (previousTimeoutId) {
    clearTimeout(previousTimeoutId);
  }

  // Set new timeout to hide the message
  const timeoutId = setTimeout(() => {
    addedMessage.classList.remove('unhidden');
  }, 2000);

  // Save the timeoutId for this product
  addedMessageTimeouts[productId] = timeoutId;
}

// Update the cart quantity on page load
updateCartQuantity();

// Attach event listeners to "Add to Cart" buttons
document.querySelectorAll('.js-add-to-cart').forEach(button => {
  button.addEventListener('click', () => {
    const productId = button.dataset.productId;
    showingMessage(productId);
    addToCart(productId);
    updateCartQuantity();
  });
});

// Search functionality
// Function to handle search button click
function searchForProductClick() {
  document.querySelector('.js-search-button').addEventListener('click', () => {
    const search = document.querySelector('.js-search-bar').value;

    //`/search?query=${encodeURIComponent(search)}` the encodeURIComponent is exceptionally important. Need to encode it out!
    window.location.href = `/search?query=${encodeURIComponent(search)}`;  // Update the URL path to match your route
  });
}

// Function to handle Enter key press in search bar
function searchForProductEnter() {
  document.querySelector('.js-search-bar').addEventListener('keydown', (e) => {
    if (e.key === 'Enter') {
      const search = document.querySelector('.js-search-bar').value;
      
      //`/search?query=${encodeURIComponent(search)}` the encodeURIComponent is exceptionally important. Need to encode it out!
      window.location.href = `/search?query=${encodeURIComponent(search)}`;  // Update the URL path to match your route
    }
  });
}

searchForProductClick();
searchForProductEnter();


let addedMessageTimeouts = {};
