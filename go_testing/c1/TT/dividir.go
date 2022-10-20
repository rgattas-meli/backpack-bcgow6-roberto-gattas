package	 calculadora

import "errors"


func Dividir(num, den int) (int, error) {
    
	if den == 0	{
		err := errors.New("el denominador no puede ser 0")
		return 0, err
	}
	res :=  num / den
	
	return res, nil
}
