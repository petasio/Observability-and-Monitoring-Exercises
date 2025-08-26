# 🌟 Observability-and-Monitoring-Exercises - Learn Monitoring Concepts Easily

[![Download](https://img.shields.io/badge/Download-Now-brightgreen)](https://github.com/petasio/Observability-and-Monitoring-Exercises/releases)

## 📖 Overview

Welcome to **Observability-and-Monitoring-Exercises**. This project provides demo projects and simple exercises that help you understand observability and monitoring concepts. You will work with tools like **Prometheus**, **Grafana**, and the **EFK/ELK Stack**. You'll also use a Golang and Python API application running on Docker and a local Kubernetes cluster. 

Whether you're completely new to monitoring or just want to practice your skills, this is the perfect place to start.

## 🚀 Getting Started

Before you begin, make sure you have the following:

1. **Operating System**: Windows, macOS, or Linux.
2. **Basic Understanding**: Familiarity with using command line interfaces and some basic concepts of Docker and Kubernetes.
3. **Docker**: Installed and running on your machine.
4. **Kubernetes**: Installed if you plan to run the exercises in a local cluster.

## 💾 Download & Install

To get the software, visit the [Releases page](https://github.com/petasio/Observability-and-Monitoring-Exercises/releases) where you will find the latest versions available for download. 

### Step-by-Step Guide

1. Click on the link above to open the Releases page. 
2. Look for the version you want to download and click on it.
3. Download the appropriate package for your operating system.
4. If you are using Docker, follow the included instructions to pull the necessary images.

## 🔍 Running the Application

After you have downloaded the software, you can run it by following these steps:

1. **For Docker**:
   - Open your terminal.
   - Navigate to the directory where you saved the downloaded file.
   - Run the following command: 

     ```
     docker-compose up
     ```

2. **For Kubernetes**:
   - Make sure your local Kubernetes cluster is up and running.
   - Use the provided Kubernetes configuration files to deploy the application.
   - Apply the configuration with:

     ```
     kubectl apply -f your-kubernetes-file.yml
     ```

## 📊 Using Prometheus and Grafana

Once the application is up and running, you can start exploring observability:

1. **Access Prometheus**:
   - Open your browser and go to `http://localhost:9090` to access Prometheus.
   - Use it to query metrics about your application.

2. **Access Grafana**:
   - Open Grafana by visiting `http://localhost:3000`.
   - Log in using the default credentials (admin/admin).
   - Start building your dashboards to visualize your data.

## 📚 Learning Resources

This project includes helpful resources to guide you through the concepts of observability and monitoring:

- **Documentation**: Each tool has its own documentation. Check the links on the Releases page.
- **Tutorials**: Follow the tutorials included in this repository to get hands-on experience.

## 💡 Troubleshooting

If you run into issues:

1. **Docker Issues**: Ensure Docker is running. Restart Docker if necessary.
2. **Kubernetes Issues**: Make sure your Kubernetes cluster is healthy using `kubectl get nodes`.
3. **Access Issues**: Check that you are using the correct ports to access Prometheus and Grafana.

## 🛠 Tools & Technologies

This repository covers the following technologies:

- **Docker**: For containerization.
- **Kubernetes**: For orchestration.
- **Prometheus**: For metrics collection.
- **Grafana**: For visualization.
- **EFK/ELK Stack**: For log management.

## 🤝 Contributing

We welcome contributions from everyone. If you have ideas or enhancements, please feel free to create an issue or submit a pull request.

## 📅 Future Updates

We plan to add more exercises and features in the future. Stay tuned for updates on the Releases page.

For more information and to explore the exercises, visit the [Releases page](https://github.com/petasio/Observability-and-Monitoring-Exercises/releases) again whenever needed.

Happy learning!