package com.steve.steve_springboot;

import jakarta.persistence.*;

@Entity
@Table(name = "cart_operator")
public class operator_cart_items {
    @Id
    @GeneratedValue(strategy = GenerationType.AUTO)
    private int id, cart_id, user_id, quantity, product_id;
    public operator_cart_items() {}
    public operator_cart_items(int cart_id, int id, int user_id, int quantity, int product_id) {
        this.cart_id = cart_id;
        this.id = id;
        this.user_id = user_id;
        this.quantity = quantity;
        this.product_id = product_id;
    }

    public int getId() {
        return id;
    }

    public void setId(int id) {
        this.id = id;
    }

    public int getCart_id() {
        return cart_id;
    }

    public void setCart_id(int cart_id) {
        this.cart_id = cart_id;
    }

    public int getUser_id() {
        return user_id;
    }

    public void setUser_id(int user_id) {
        this.user_id = user_id;
    }

    public int getQuantity() {
        return quantity;
    }

    public void setQuantity(int quantity) {
        this.quantity = quantity;
    }

    public int getProduct_id() {
        return product_id;
    }

    public void setProduct_id(int product_id) {
        this.product_id = product_id;
    }
}
