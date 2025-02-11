var cart = [];
function findAmount() {
    let qty = parseInt(document.getElementById("qty").value);
    let price = parseFloat(document.getElementById("price").value);
    let amount = qty * price;
    document.getElementById("amount").value = amount;
}
function addItem() {
    let name = document.getElementById("name").value;
    let qty = parseInt(document.getElementById("qty").value);
    let price = parseFloat(document.getElementById("price").value);
    let amount = parseFloat(document.getElementById("amount").value);

    // duplicate handling
    let productExist = cart.find(someProduct => someProduct.name === name);
    
    /*
    writing with reduce function
    let productExist = cart.reduce((found, someProduct) => {
        return someProduct.name === name? someProduct : found;
    }, null);
    */

    if(productExist){
        productExist.qty += qty;
        productExist.amount += amount;
    }
    else {
        let product = {name:name, qty:qty, price:price, amount:amount};
        cart.push(product);
    }

    // clear boxes
    document.getElementById("name").value = ""
    document.getElementById("qty").value = ""
    document.getElementById("price").value = ""
    document.getElementById("amount").value = ""
    
    // cart html
    let cart_rows = "";
    for(let prod of cart) {
        cart_rows = cart_rows + `<TR><TD>${prod.name}</TD>
                <TD>${prod.qty}</TD>
                <TD>${prod.price}</TD>
                <TD>${prod.amount}</TD>
                </TR>`;
    }
    let totalAmount = cart.reduce(function(sum,prod){
        return sum + parseFloat(prod.amount);
    },0)
    let table = `<TABLE class="table table-border">
        <TR>    
            <TH>PRODUCT</TH>
            <TH>QTY</TH>
            <TH>PRICE</TH>
            <TH>AMOUNT</TH>       
        </TR>
        ${cart_rows}
        <TR>    
            <TH colspan="3">Total</TH>
            <TH>${totalAmount}</TH>       
        </TR>
        </TABLE>`
    document.getElementById("scart").innerHTML = table;
}