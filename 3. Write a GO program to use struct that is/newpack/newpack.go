package newpack 
type Number struct { 
    Value int 
} 
func (n Number) Square() int { 
    return n.Value * n.Value 
} 
