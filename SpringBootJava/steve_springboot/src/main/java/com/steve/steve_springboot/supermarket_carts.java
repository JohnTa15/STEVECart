package com.steve.steve_springboot;

import jakarta.persistence.*;

@Entity
@Table(name = "carts")
public class supermarket_carts {
    @Id
    @GeneratedValue(strategy = GenerationType.AUTO)
    private int id;
    private boolean isActive;
    private String fw_version;

    public supermarket_carts() {}
    public supermarket_carts(int id, boolean isActive, String fw_version) {
        this.id = id;
        this.isActive = isActive;
        this.fw_version = fw_version;
    }

    public boolean isIsActive() {
        return isActive;
    }

    public void setIsActive(boolean isActive) {
        this.isActive = isActive;
    }

    public int getId() {
        return id;
    }

    public void setId(int id) {
        this.id = id;
    }

    public String getFw_version() {
        return fw_version;
    }

    public void setFw_version(String fw_version) {
        this.fw_version = fw_version;
    }
}
