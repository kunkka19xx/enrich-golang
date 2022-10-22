package noah.mapper;

import java.util.ArrayList;

import org.apache.ibatis.annotations.Mapper;
import org.springframework.stereotype.Repository;

import noah.model.Product;

@Mapper
@Repository
public interface ProductMapper {
	boolean insertProduct(Product product);

	ArrayList<Product> getAllProducts();
	
	
	boolean updateProductByName(Product product);

	boolean deleteProduct(long id);
}
