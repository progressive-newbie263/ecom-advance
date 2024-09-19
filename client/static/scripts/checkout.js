import { renderCheckoutHeader } from "./checkout/checkoutHeader.js";
import { renderOrderSummary } from "./checkout/orderSummary.js";
import { renderPaymentSummary } from "./checkout/paymentSummary.js";
import { loadCart, loadCartFetch } from "../data/cart.js";

async function loadPage() {
 
  try {
    await Promise.all([
      loadCartFetch()
    ]);
  } catch(error) {
    console.log('Unexpected Error. Please try again later.');
  };
 
  renderCheckoutHeader();
  renderOrderSummary();
  renderPaymentSummary();
}





