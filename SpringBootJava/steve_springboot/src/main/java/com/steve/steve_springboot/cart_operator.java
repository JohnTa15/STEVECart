package com.steve.steve_springboot;

import java.time.LocalDateTime;

import jakarta.persistence.*;

@Entity
@Table(name = "cart_operator")
public class cart_operator {
    @Id
    @GeneratedValue(strategy = GenerationType.AUTO)
    private int id, cart_id, user_id;
    private LocalDateTime event_time;
    public cart_operator() {}
    public cart_operator(int cart_id, int id, LocalDateTime event_time, int user_id) {
        this.cart_id = cart_id;
        this.id = id;
        this.event_time = event_time;
        this.user_id = user_id;
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

    public LocalDateTime getEvent_time() {
        return event_time;
    }

    public void setEvent_time(LocalDateTime event_time) {
        this.event_time = event_time;
    }
}
