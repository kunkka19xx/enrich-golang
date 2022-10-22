package noah.service;

import java.util.ArrayList;

import org.springframework.stereotype.Service;

import noah.model.Product;

@Service
public interface ProductService {
	ArrayList<Product> getAllProducts();

	boolean insertProduct(Product product);
	
	boolean updateProductByName(Product product);
	
	boolean deleteProduct(long id);
}
