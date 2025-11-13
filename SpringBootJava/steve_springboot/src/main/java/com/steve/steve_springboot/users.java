package com.steve.steve_springboot;

import java.time.LocalDateTime;

import jakarta.persistence.*;

@Entity
@Table(name = "users")
public class users {
    @Id
    @GeneratedValue(strategy = GenerationType.AUTO)
    private int id;
    private String username, email, password_hash;
    private LocalDateTime created_user_at;
    public users() {}
    public users(int id, String email, String password_hash, LocalDateTime created_user_at, String username) {
        this.id = id;
        this.email = email;
        this.password_hash = password_hash;
        this.created_user_at = created_user_at;
        this.username = username;
    }
    public String getEmail() {
        return email;
    }

    public void setEmail(String email) {
        this.email = email;
    }

    public String getPassword_hash() {
        return password_hash;
    }

    public void setPassword_hash(String password_hash) {
        this.password_hash = password_hash;
    }

    public String getUsername() {
        return username;
    }

    public void setUsername(String username) {
        this.username = username;
    }

    public LocalDateTime getCreated_user_at() {
        return created_user_at;
    }

    public void setCreated_user_at(LocalDateTime created_user_at) {
        this.created_user_at = created_user_at;
    }

    public int getId() {
        return id;
    }

    public void setId(int id) {
        this.id = id;
    }
}
