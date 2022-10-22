package noah.controller;

import java.util.ArrayList;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.DeleteMapping;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.PutMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import noah.model.Product;
import noah.service.ProductService;

@RestController()
@RequestMapping(path = "/product")
public class ProductController {
	@Autowired
	private ProductService productService;

	@PostMapping("/insertion")
	private boolean insertProduct(@RequestBody Product product) {
		return productService.insertProduct(product);
	}

	@GetMapping
	private ArrayList<Product> getAllProducts() {
		return productService.getAllProducts();
	}

	@PutMapping
	private Product updateProductByName(@RequestBody Product product) {
		if (productService.updateProductByName(product))
			return product;
		else
			return null;
	}
	
	@DeleteMapping(value = "/{id}")
	private boolean deleteProduct(@PathVariable(name = "id") long id) {
		return productService.deleteProduct(id);
	}
}
