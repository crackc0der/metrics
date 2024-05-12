RUN: make up

**Categoryis**:
    GETALL:
        GET: get/categories

    ADD:
        {
            "CategoryName: "category_name"
        }

        GET: add/category

    UPDATE:
        {
            "Id": category_id
        }

        GET: update/category

    DELETE:
        {
            "Id": category_id
        }

        GET: delete/category

**Peoducts**:
    GETALL:
        GET: get/products

    GET:
        {
            "Id": product_id
        }

        GET: get/product

    ADD:
        {
            ProductName": "product_name",
            "ProductCategory": product_category_id,
            "ProductPrice": peoduct_price
        }

        GET: add/product

    UPDATE:
        {
            "Id": product_id,
            "ProductName": "product_name",
            "ProductCategory": product_category_id,
            "ProductPrice": peoduct_price            
        }

        GET: update/product

    DELETE:
        {
            "ProductId": product_id
        }

        GET: delete/product