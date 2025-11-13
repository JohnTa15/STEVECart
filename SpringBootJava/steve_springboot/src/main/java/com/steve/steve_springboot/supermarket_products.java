package com.steve.steve_springboot;

import jakarta.persistence.*;

@Entity
@Table(name = "products")

public class supermarket_products {
    @Id
    @GeneratedValue(strategy = GenerationType.AUTO)
    private int id, pcs, day, month, year;
    private String name, category,description;
    private Double price;
    public supermarket_products() {}
    public supermarket_products(String category, int day, String description, int id, int month, String name, int pcs, Double price, int year) {
        this.category = category;
        this.day = day;
        this.description = description;
        this.id = id;
        this.month = month;
        this.name = name;
        this.pcs = pcs;
        this.price = price;
        this.year = year;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public String getCategory() {
        return category;
    }

    public void setCategory(String category) {
        this.category = category;
    }

    public String getDescription() {
        return description;
    }

    public void setDescription(String description) {
        this.description = description;
    }
}
