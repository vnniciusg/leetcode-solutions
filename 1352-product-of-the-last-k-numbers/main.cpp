class ProductOfNumbers {
public:
    vector<int> prefix_product;

    ProductOfNumbers() {
        prefix_product.push_back(1);
    }
    
    void add(int num) {
        if (num == 0) {
            prefix_product.clear();
            prefix_product.push_back(1);
        }else {
            prefix_product.push_back(prefix_product.back() * num);
        }
    }
    
    int getProduct(int k) {
        if (k >= prefix_product.size()) {
            return 0;
        }

        return prefix_product.back() / prefix_product[prefix_product.size() - 1 - k];
    }
};
