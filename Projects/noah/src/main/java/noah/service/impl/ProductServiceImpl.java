package noah.service.impl;

import java.util.ArrayList;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import noah.mapper.ProductMapper;
import noah.model.Product;
import noah.service.ProductService;

@Service
public class ProductServiceImpl implements ProductService{
	
	@Autowired
	ProductMapper productMapper;
	@Override
	public boolean insertProduct(Product product) {
		return productMapper.insertProduct(product);
	}
	@Override
	public ArrayList<Product> getAllProducts() {
		return productMapper.getAllProducts();
	}
	@Override
	public boolean updateProductByName(Product product) {
		return productMapper.updateProductByName(product);
	}
	@Override
	public boolean deleteProduct(long id) {
		return productMapper.deleteProduct(id);
	}

}
