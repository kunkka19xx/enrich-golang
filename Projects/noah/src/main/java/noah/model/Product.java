package noah.model;

import lombok.Data;

@Data
public class Product {
	private long id;
	
	private String name;
	
	private String newName;
	
	private String oldName;
}
