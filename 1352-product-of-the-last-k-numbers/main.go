type ProductOfNumbers struct {
    prefixProduct []int
    size int
}


func Constructor() ProductOfNumbers {
    return ProductOfNumbers{
        prefixProduct: []int{1},
        size: 0,
    }
}


func (this *ProductOfNumbers) Add(num int)  {
    if num == 0 {
        this.prefixProduct = []int{1}
        this.size = 0
    }else {
        this.prefixProduct = append(this.prefixProduct, this.prefixProduct[this.size] * num)
        this.size++
    }
}


func (this *ProductOfNumbers) GetProduct(k int) int {
    if k > this.size {
        return 0
    }

    return this.prefixProduct[this.size] / this.prefixProduct[this.size - k]
}
