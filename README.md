# leetcode\_go

一个使用 Go（`go.mod`）进行 LeetCode 练习题与 Weekly 周赛题整理的代码仓库。内容以「可运行」「可复用」「按专题归档」为目标，包含常见题型的模板、数据结构实现与周赛解题记录。

---

## 1\. 环境与约定

- 语言：Go
- 依赖管理：Go Modules（见 `go.mod`）
- IDE：GoLand
- 运行方式：以 `main.go` 为入口组织每个专题/周赛的样例与调试代码（偏向练习与快速验证）

> 说明：本仓库偏「练习场」风格，并非单一可执行程序。不同目录下可能存在多个 `main.go` 入口文件，用于按专题或按周赛独立运行。

---

## 2\. 项目目录结构

项目根目录 `leetcode_go` 下主要分为以下几块：

### 2\.1 `cmd/`
- `cmd/main/main.go`  
  可作为统一入口（如果你希望后续把各专题/周赛的调用集中到这里，也适合放汇总调用逻辑）。

### 2\.2 `internal/`
用于存放按「周赛」或「专题」拆分的练习代码（内部包，不建议被外部项目直接依赖）。

#### 2\.2\.1 周赛：`internal/exercise/`
- `internal/exercise/weekly429/main.go`
- `internal/exercise/weekly475/main.go`
- `internal/exercise/weekly476/main.go`
- `internal/exercise/weekly480/main.go`

每个周赛目录通常是一个独立的小练习单元：包含当周赛题的解题实现、调试入口、必要的辅助函数等。你可以在对应 `main.go` 里快速跑样例、打印中间结果、验证边界条件。

#### 2\.2\.2 Hot 100：`internal/hot100/`
用于整理高频经典题/常见面试题集合，按题型拆分：

- `internal/hot100/main.go`
- `internal/hot100/doubleptr/main.go`：双指针
- `internal/hot100/dp/main.go`：动态规划
- `internal/hot100/greedy/main.go`：贪心
- `internal/hot100/hash/main.go`：哈希/映射
- `internal/hot100/linkedlist/main.go`：链表
- `internal/hot100/multidp/main.go`：多维 DP / 状态压缩等

> 目标：把同类题的套路沉淀成可复用的思路与模板，例如：滑动窗口、前缀和、二分、区间 DP、单调栈等。

#### 2\.2\.3 LeetCode 75：`internal/leetcode75/`
按官方/社区常用的 75 题单归档（同样按题型拆分）：

- `internal/leetcode75/main.go`
- `internal/leetcode75/binary_search/main.go`：二分
- `internal/leetcode75/bit_operation/main.go`：位运算
- `internal/leetcode75/depth_first_search/main.go`：DFS/回溯
- `internal/leetcode75/dynamic_programming/main.go`：DP
- `internal/leetcode75/multi_dynamic_programming/main.go`：多维 DP
- `internal/leetcode75/slide_window/main.go`：滑动窗口

> 说明：目录中存在 `main.go~` 这类文件通常是编辑器备份文件，可忽略或后续清理。

#### 2\.2\.4 随机练习/数据结构：`internal/random/`
更偏向「按数据结构/实现」的练习区域，包含题解与手写结构体：

- `internal/random/array/main.go`：数组相关
- `internal/random/string/main.go`：字符串相关
- `internal/random/hashmap/main.go`：哈希
- `internal/random/btree/`：二叉树相关（含 `type.go`）
- `internal/random/btree_fill/node.go`：树节点或填充类题型
- `internal/random/linkedlist/`：链表相关（包含多种实现）
    - `circularlinkedlist/design_linked_list.go`
    - `singlelinkedlist/design_linked_list.go`
    - `lru/main.go`：LRU 实现/练习
    - `lfu/main.go`：LFU 实现/练习
- `internal/random/stack/`：栈/队列相关
    - `stack.go` / `queue.go`
    - `mock_queue/queue.go`、`mock_stack/stack.go`、`mock_stack_adv/stack.go`
    - `main/main.go`：可运行入口
- `internal/types/`：类型/基础结构练习
    - `binary_search/main.go`：二分相关习题或模板

---

## 3\. 公共数据结构：`pkg/`

`pkg/` 用于放置可复用的数据结构定义，供各题解引用：

- `pkg/list_node.go`：链表节点（ListNode）
- `pkg/tree_node.go`：二叉树节点（TreeNode）

> 建议：题解中尽量复用 `pkg/` 的节点类型，减少重复定义，保证不同专题可以共享工具函数（如建树、打印链表等）。

---

## 4\. 如何运行

由于每个目录下可能都有独立的 `main.go`，推荐在项目根目录直接运行指定包（以 macOS + Go Modules 为例）：

- 运行某个周赛：
    - `go run ./internal/exercise/weekly480`

- 运行某个专题（例如 Hot100 的 dp）：
    - `go run ./internal/hot100/dp`

- 运行 LeetCode 75 的滑动窗口：
    - `go run ./internal/leetcode75/slide_window`

- 运行 `cmd` 入口：
    - `go run ./cmd/main`

> 在 GoLand 中也可以直接对某个目录下的 `main.go` 创建 Run Configuration，按包运行更方便。

---

## 5\. 代码组织建议（本仓库适用的约定）

为保证题目数量增多后仍易维护，建议每个专题/周赛目录遵循类似结构（当前以 `main.go` 为主也完全可行）：

- `main.go`
    - `func main()`：只放快速验证用例或手动调试入口
    - 题解函数：以题目名/题号命名（例如 `func maxSubArray(nums []int) int`）
    - 辅助函数：如 `min/max/abs`、打印函数、初始化数据等
- 若某一目录题目变多，可拆成多文件：
    - `solution_XXXX.go`
    - `utils.go`

命名推荐：
- 同一目录内按题型分组或按题号排序
- 尽量保持函数签名与 LeetCode 原题一致，便于复制提交

---

## 6\. 内容覆盖范围

本仓库主要覆盖：
- 高频题：数组、字符串、哈希、双指针、滑动窗口、二分
- 数据结构：链表、栈、队列、堆（如后续加入）、二叉树/DFS/BFS
- 算法思想：贪心、动态规划（含多维 DP）、位运算
- 系统设计类实现题：`LRU`、`LFU`、设计链表等

---

## 7\. 练习与复盘方式（建议流程）

1. 在对应专题目录下实现题解函数
2. 在 `main()` 中加入最小可验证样例（正常/边界/性能）
3. 通过 `go run` 运行目录包验证
4. 周赛题则按场次目录沉淀，后续回看方便对比思路与常见坑点

---

## 8\. 备注

- `README.md` 当前由项目结构自动补全而来，后续可按个人习惯补充：
    - 每个专题包含的题单链接
    - 常用模板（滑动窗口、二分上下界、前缀和、DP 初始化等）
    - 常见 bug 记录（越界、溢出、边界条件、复杂度超限）

---
