db = db.getSiblingDB('db_aapanavypar');

db.createCollection('userData');
db.createCollection('orderData');
db.createCollection('shopData');
db.createCollection('productData');
db.createCollection('analyticalData');


/*

For Shop
*
* shopId TEXT NOSTEM NOINDEX
* shopName TEXT
* primaryImage TEXT NOSTEM NOINDEX
* categoryOfShop TEXT NOSTEM SORTABLE
* ratingOfShop NUMERIC SORTABLE
* shopkeeper TEXT NOSTEM NOINDEX
* location GEO SORTABLE
*


For Trending search we can find shop near by and then find the products of that shop.
For Search We Can store only the product-id and category and name and then perform search on it.


For Product
*
* productId TEXT NOSTEM NOINDEX
* productName TEXT
* categoryOfProduct TEXT NOSTEM SORTABLE
* likesOfProduct NUMERIC SORTABLE
*
*/
