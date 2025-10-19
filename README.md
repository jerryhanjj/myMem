# MyMem - 时间记录器

一个精美的时间记录器应用，支持倒计时和累计天数功能，使用 Vue 3 + Element Plus 前端和 Go + SQLite 后端开发。

## ✨ 功能特性

- 📅 **倒计时模式**：记录未来事件，实时显示剩余天数/时间
- ⏱️ **累计天数模式**：记录过去事件，显示已经过的天数
- 🎨 **美观的卡片式界面**：使用 Element Plus UI 框架，界面精美
- 🎯 **实时更新**：时间自动刷新，无需手动操作
- 🏷️ **分类管理**：支持为记录添加分类标签
- 🎨 **自定义颜色**：每个记录可以设置独特的主题颜色
- 🔍 **搜索筛选**：快速查找特定记录
- 📱 **响应式设计**：完美适配桌面和移动设备
- 🎭 **多种显示模式**：支持简洁、详细、精确、智能四种时间显示方式

## 🛠️ 技术栈

### 前端
- **框架**：Vue 3 (Composition API)
- **UI 库**：Element Plus
- **状态管理**：Pinia
- **HTTP 客户端**：Axios
- **构建工具**：Vite
- **日期处理**：Day.js

### 后端
- **语言**：Go 1.21+
- **Web 框架**：Gin
- **ORM**：GORM
- **数据库**：SQLite

## 📦 项目结构

```
myMem/
├── backend/                    # Go 后端
│   ├── database/              # 数据库连接
│   │   └── database.go
│   ├── handlers/              # API 处理器
│   │   └── time_record_handler.go
│   ├── models/                # 数据模型
│   │   └── time_record.go
│   ├── main.go                # 入口文件
│   ├── go.mod                 # Go 依赖
│   └── .gitignore
│
├── frontend/                   # Vue 前端
│   ├── src/
│   │   ├── api/               # API 接口
│   │   │   └── index.js
│   │   ├── components/        # Vue 组件
│   │   │   ├── TimeCard.vue   # 时间卡片
│   │   │   └── TimeForm.vue   # 表单组件
│   │   ├── stores/            # Pinia 状态
│   │   │   └── timeRecord.js
│   │   ├── views/             # 页面视图
│   │   │   └── Home.vue
│   │   ├── App.vue            # 根组件
│   │   ├── main.js            # 入口文件
│   │   └── style.css          # 全局样式
│   ├── index.html
│   ├── package.json
│   ├── vite.config.js
│   └── .gitignore
│
└── README.md
```

## 🚀 快速开始

### 前置要求

- **Node.js**: 18.0+ 和 npm
- **Go**: 1.21+
- **Git**: 用于版本控制

### 安装步骤

#### 1. 克隆项目（如果需要）

```bash
cd myMem
```

#### 2. 启动后端服务

```bash
# 进入后端目录
cd backend

# 安装 Go 依赖
go mod tidy

# 运行后端服务
go run main.go
```

后端服务将在 `http://localhost:8080` 启动

#### 3. 启动前端应用

打开新的终端窗口：

```bash
# 进入前端目录
cd frontend

# 安装依赖
npm install

# 启动开发服务器
npm run dev
```

前端应用将在 `http://localhost:5173` 启动

#### 4. 访问应用

打开浏览器访问：`http://localhost:5173`

## 📖 使用指南

### 创建时间记录

1. 点击页面右上角的 **"创建记录"** 按钮
2. 填写以下信息：
   - **标题**：记录的名称（必填）
   - **描述**：详细说明（可选）
   - **目标日期**：选择日期和时间（必填）
   - **记录类型**：选择 "倒计时" 或 "累计天数"
   - **分类**：自定义分类标签（可选）
   - **颜色**：选择卡片主题颜色
3. 点击 **"创建"** 按钮

### 编辑记录

- 点击卡片右上角的 **编辑图标**
- 修改信息后点击 **"更新"**

### 删除记录

- 点击卡片右上角的 **删除图标**
- 确认删除操作

### 筛选和搜索

- 使用顶部的 **按钮组** 筛选不同类型的记录
- 使用 **搜索框** 根据标题或分类搜索

## 🎯 API 接口

后端提供以下 RESTful API：

| 方法 | 路径 | 描述 |
|------|------|------|
| GET | `/api/records` | 获取所有记录 |
| GET | `/api/records/:id` | 获取单个记录 |
| POST | `/api/records` | 创建新记录 |
| PUT | `/api/records/:id` | 更新记录 |
| DELETE | `/api/records/:id` | 删除记录 |
| GET | `/health` | 健康检查 |

### 请求示例

创建记录：

```bash
curl -X POST http://localhost:8080/api/records \
  -H "Content-Type: application/json" \
  -d '{
    "title": "新年倒计时",
    "description": "2026年的到来",
    "target_date": "2026-01-01T00:00:00.000Z",
    "record_type": "countdown",
    "category": "节日",
    "color": "#ff6b6b"
  }'
```

## 🗄️ 数据库

SQLite 数据库会自动在后端目录创建（`mymem.db`），包含以下表结构：

### time_records 表

| 字段 | 类型 | 说明 |
|------|------|------|
| id | INTEGER | 主键，自增 |
| title | TEXT | 标题 |
| description | TEXT | 描述 |
| target_date | DATETIME | 目标日期 |
| record_type | TEXT | 类型（countdown/elapsed） |
| category | TEXT | 分类 |
| color | TEXT | 颜色 |
| created_at | DATETIME | 创建时间 |
| updated_at | DATETIME | 更新时间 |

## 🎨 界面预览

- **渐变背景**：紫色渐变背景，视觉效果优雅
- **玻璃拟态**：半透明卡片设计，现代感十足
- **实时动画**：时间数字实时跳动更新
- **悬停效果**：卡片悬停时有平滑的上浮效果
- **响应式布局**：自动适配不同屏幕尺寸

## 🔧 开发命令

### 后端

```bash
# 运行开发服务器
go run main.go

# 构建可执行文件
go build -o mymem-server main.go

# 运行测试
go test ./...
```

### 前端

```bash
# 安装依赖
npm install

# 开发模式
npm run dev

# 构建生产版本
npm run build

# 预览生产构建
npm run preview
```

## 📝 配置说明

### 后端配置

在 `backend/main.go` 中可以修改：
- 服务器端口（默认 8080）
- CORS 允许的源
- 数据库路径

### 前端配置

在 `frontend/vite.config.js` 中可以修改：
- 开发服务器端口（默认 5173）
- API 代理设置
- 构建输出路径

## 🚢 部署

### 后端部署

1. 构建可执行文件：
   ```bash
   cd backend
   go build -o mymem-server main.go
   ```

2. 运行服务器：
   ```bash
   ./mymem-server
   ```

### 前端部署

1. 构建生产版本：
   ```bash
   cd frontend
   npm run build
   ```

2. 将 `dist` 目录部署到任何静态服务器（Nginx、Apache 等）

## 🤝 贡献

欢迎提交 Issue 和 Pull Request！

## 📄 许可证

MIT License

## 👨‍💻 作者

开发者：MyMem Team

---

**享受记录时间的乐趣！** 🎉
