# AapanaVypar ViewProvider Service #

## RPCs :

- **GetTrendingShops**
- **GetTrendingProductsByShop**

- **GetProduct**
- **GetShop**

- **GetProductsBySearch**
- **GetShopsBySearch**

- **AddToLikeProduct**
- **RemoveFromLikeProduct**

- **AddToCartProduct**
- **RemoveFromCartProduct**

- **GetOrders**
- **GetCart**

- **GetProfile**
- **UpdateAddress**

- **InitUser**

### GetTrendingShops

Request

    message GetTrendingShopsRequest {
        string apiKey = 10;
        string token = 11;
        Location location = 12;
        string distanceInMeter = 13; 
    }

Response

    message GetTrendingShopsResponse {
        ShopsNearBy shops = 14;
    }

### GetTrendingProductsByShop

Request

    message GetTrendingProductsByShopRequest {
        string apiKey = 21;
        string token = 22;
        repeated string shopId = 23;
    }

Response

    message GetTrendingProductsByShopResponse {
        ProductsOfShopsNearBy categoryData = 24;
    }

### GetProduct

Request

    message GetProductRequest {
        string apiKey = 41;
        string token = 42;
        string productId = 43;
        string shopId = 44;
    }

Response

    message GetProductResponse {
        string productId = 45;
        string shopId = 46;
        string shopName = 47;
        string productName = 48;
        string productDescription = 49;
        string shippingInfo = 50;
        uint32 stock = 51;
        uint64 likes = 52;
        float price = 53;
        uint32 offer = 54;
        repeated string images = 55;
        repeated Category category = 56;
        string timestamp = 57; 
    }

### GetProductsBySearch

Request

    message GetProductsBySearchRequest {
        string apiKey = 85;
        string token = 86;
        string search = 87;
        repeated string shopIds = 88;
    }

Response

    message GetProductsBySearchResponse {
        ProductsOfShopsNearBy products = 89;
    }

### GetShopsBySearch

Request

    message GetShopsBySearchRequest {
        string apiKey = 90;
        string token = 91;
        string search = 92;
        string distanceInMeter = 93;
        Location location = 94;
    }


Response

    message GetShopsBySearchResponse {
        ShopsNearBy shops = 95;
    }

### AddToLikeProduct

Request

    message AddToLikeProductRequest {
        string token = 25;
        string apiKey = 26;
        string productId = 27;
    }

Response

    message AddToLikeProductResponse {
        bool status = 28;
    }

### RemoveFromLikeProduct

Request

    message RemoveFromLikeProductRequest {
        string token = 29;
        string apiKey = 30;
        string productId = 31;
    }

Response

    message RemoveFromLikeProductResponse {
        bool status = 32;
    }

### AddToCartProduct

Request

    message AddToCartProductRequest {
        string token = 33;
        string apiKey = 34;
        string productId = 35;
    }

Response

    message AddToCartProductResponse {
        bool status = 36;
    }

### RemoveFromCartProduct

Request

    message RemoveFromCartProductRequest {
        string token = 37;
        string apiKey = 38;
        string productId = 39;
    }

Response

    message RemoveFromCartProductResponse {
        bool status = 40;
    }

### GetOrders

Request

    message GetOrdersRequest {
        string token = 96;
        string apiKey = 97;
    }

Response

    message GetOrdersResponse {
        string orderId = 98;
        Status status = 99;
        string productId = 100;
        string deliveryTimeStamp = 101;
        string orderTimeStamp = 102;
        float price = 103;
        uint32 quantity = 104;
        string productName = 105;
        string productImage = 106;
    }

### GetCart

Request

    message GetCartRequest {
        string token = 107;
        string apiKey = 108;
    }

Response

    message GetCartResponse {
        ProductsOfShopsNearBy products = 109;
    }

### GetProfile

Request

    message GetProfileRequest {
        string token = 110;
        string apiKey = 111;
    }

Response

    message GetProfileResponse {
        string userName = 112;
        Address address = 113;
    }

### UpdateAddress

Request

    message UpdateAddressRequest {
        string token = 114;
        string apiKey = 115;
        Address address = 116;
    }

Response

    message UpdateAddressResponse {
        bool status = 117;
    }

### InitUser

Request

    message InitUserRequest {
        string token = 127;
        string apiKey = 128;
    }

Response

    message InitUserResponse {
        bool status = 129;
    }

## Error Codes

### 1. GetTrendingShops :

- **Unauthenticated** : No API Key Is Specified  
- **Unauthenticated** : Request With Invalid Token
- **InvalidArgument** : Invalid Location
- **InvalidArgument** : Invalid Distance
- **NotFound**        : Unable To Get Data For Shop
- **Internal**        : Unable To Parse Data
- **Unknown**         : Stream Error

### 2. GetTrendingProductsByShop :

- **Unauthenticated** : No API Key Is Specified  
- **Unauthenticated** : Request With Invalid Token
- **Unknown**         : Unable To Provide Data For Given Shops
- **Internal**        : Unable To Parse Data
- **Unknown**        : Stream Error

### 3. GetProduct :

- **Unauthenticated** : No API Key Is Specified  
- **Unauthenticated** : Request With Invalid Token
- **InvalidArgument** : Invalid Product Id
- **NotFound**        : Unable To Get Product

### 4. GetShop :

- **Unauthenticated** : No API Key Is Specified  
- **Unauthenticated** : Request With Invalid Token
- **NotFound**        : Unable To Get Shop

### 5. GetProductsBySearch :

- **Unauthenticated** : No API Key Is Specified  
- **Unauthenticated** : Request With Invalid Token
- **Internal**        : Problem Occurred
- **Internal**        : Unable To Parse Data
- **NotFound**        : Unable To Get Data

### 6. GetShopsBySearch :

- **Unauthenticated** : No API Key Is Specified
- **Unauthenticated** : Request With Invalid Token
- **Internal**        : Problem Occurred
- **InvalidArgument** : Invalid Location
- **InvalidArgument** : Invalid Distance
- **Internal**        : Unable To Parse Data
- **Unknown**         : Stream Error
- **Internal**        : Error While Searching

### 7. AddToLikeProduct :

- **Unauthenticated** : No API Key Is Specified
- **Unauthenticated** : Request With Invalid Token
- **NotFound**        : Product Does Not Exist
- **Internal**        : Unable To Add Like
- **Internal**        : Unable To Add Like To Product

### 8. RemoveFromLikeProduct :

- **Unauthenticated** : No API Key Is Specified
- **Unauthenticated** : Request With Invalid Token
- **NotFound**        : Product Does Not Exist
- **Internal**        : Unable To Parse Data
- **Internal**        : Unable To Add Like To Product
- **Internal**        : Unable To UnLike

### 9. AddToCartProduct :

- **Unauthenticated** : No API Key Is Specified
- **Unauthenticated** : Request With Invalid Token
- **NotFound**        : Product Does Not Exist
- **Internal**        : Unable To Add Product To Cart

### 10. AddToCartProduct :

- **Unauthenticated** : No API Key Is Specified
- **Unauthenticated** : Request With Invalid Token
- **NotFound**        : Product Does Not Exist
- **Internal**        : Unable To Add Product To Cart

### 11. RemoveFromCartProduct :

- **Unauthenticated** : No API Key Is Specified
- **Unauthenticated** : Request With Invalid Token
- **Internal**        : Unable To Remove Product From Cart

### 12. GetOrders :

- **Unauthenticated** : No API Key Is Specified
- **Unauthenticated** : Request With Invalid Token
- **NotFound**        : Product Not Found In Cash
- **Unknown**         : Stream Error
- **Unknown**         : Error While Sending Data

### 13. GetCart :

- **Unauthenticated** : No API Key Is Specified
- **Unauthenticated** : Request With Invalid Token
- **NotFound**        : Unable To Get Cart
- **Internal**        : Unable To Get Product Info
- **Internal**        : Unable To Parse Data
- **Unknown**         : Stream Error

### 14. GetProfile :

- **Unauthenticated** : No API Key Is Specified
- **Unauthenticated** : Request With Invalid Token
- **NotFound**        : Data Not Found

### 15. UpdateAddress :

- **Unauthenticated** : No API Key Is Specified
- **Unauthenticated** : Request With Invalid Token
- **Unknown**         : Unable To Update Address

### 15. InitUser :

- **Unauthenticated** : No API Key Is Specified
- **Unauthenticated** : Request With Invalid Token
- **Unknown**         : Unable To Init User
