# ⚙️ SkyShift

SkyShift is an image transformation system -- a powerful, modular interface for dynamic, real-time visual manipulation

> Image transformation in the air — via API.

## 🌐 API Overview

## 🛠️ Effects Built-In

- _Resize_
- _Sharpen_
- _Blurr_
- _Brightness_

## 📂 Project Structure

### ./app

**Folder with business logic only**. This directory doesn't care about _what database driver you're using_ or _which caching solution your choose_ or any third-party things.

- `./app/controllers` folder for functional controllers (used in routes)
- `./app/models` folder for describe business models and methods of your project

### ./docs

**Folder with API Documentation**. This directory contains config files for auto-generated API Docs by Swagger.

### ./pkg

**Folder with project-specific functionality**. This directory contains all the project-specific code tailored only for your business use case, like _configs_, _middleware_, _routes_ or _utils_.

- `./pkg/configs` folder for configuration functions
- `./pkg/middleware` folder for add middleware (Fiber built-in and yours)
- `./pkg/routes` folder for describe routes of your project
- `./pkg/utils` folder with utility functions (server starter, error checker, etc)

### ./platform

**Folder with platform-level logic**. This directory contains all the platform-level logic that will build up the actual project, like _setting up the database_ and _cache server instance_.

- `./platform/cache` folder with in-memory cache setup functions (by default, Redis)
- `./platform/database` folder with database setup functions (by default, PostgreSQL)
