class ProductOfNumbers {
public:
    vector<int> prefix_product;
    int size;

    ProductOfNumbers() {
        prefix_product.push_back(1);
        size = 0;
    }
    
    void add(int num) {
        if (num == 0) {
            prefix_product.clear();
            prefix_product.push_back(1);
            size = 0;
        }else {
            prefix_product.push_back(prefix_product.back() * num);
            size++;
        }
    }
    
    int getProduct(int k) {
        if (k > size) {
            return 0;
        }

        return prefix_product.back() / prefix_product[prefix_product.size() - 1 - k];
    }
};

/**
 * Your ProductOfNumbers object will be instantiated and called as such:
 * ProductOfNumbers* obj = new ProductOfNumbers();
 * obj->add(num);
 * int param_2 = obj->getProduct(k);
 */
