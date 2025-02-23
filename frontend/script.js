const API_GATEWAY = 'http://192.168.1.63:8080/cart';

function getProductIdFromURL() {
    const match = window.location.pathname.match(/\/add\/(\w+)/);
    return match ? match[1] : null;
}

async function addProductToCart(productId) {
    if (!productId) return;

    const response = await fetch(`${API_GATEWAY}/add/${productId}`, {
        credentials: "include",
    });

    if (response.ok) {
        console.log(`Product ${productId} added to cart`);
    } else {
        console.error("Error adding product.");
    }

    window.location.href = "/";
}

async function addProduct() {
    const productId = document.getElementById("product_id").value.trim();
    if (!productId) {
        alert("Please, add product ID.");
        return;
    }

    addProductToCart(productId);
}


async function fetchCart() {
    const response = await fetch(API_GATEWAY, {
        credentials: "include",
    });

    if (response.ok) {
        const data = await response.json();
        return data.cart;
    }
    return [];
}

async function showCart() {
    const cartDiv = document.getElementById("cart");
    const cart = await fetchCart();

    if (cart.length === 0) {
        cartDiv.innerHTML = "<p>Cart is empty.</p>";
    } else {
        cartDiv.innerHTML = `<ul>${cart.map(item => `<li>${item}</li>`).join("")}</ul>`;
    }
}


async function init() {
    const productId = getProductIdFromURL();

    if (productId) {
        await addProductToCart(productId);
    } else {
        await showCart();
    }
}

document.addEventListener("DOMContentLoaded", init);
