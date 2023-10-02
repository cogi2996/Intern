package resources

type Response struct {
	Reason string      `json:"reason,omitempty"`
	Data   interface{} `json:"data,omitempty"`
	Error  interface{} `json:"error,omitempty"`
}

type Empty struct {
}

type CreateResponse struct {
	ID   int64  `json:"id"`
	UUID string `json:"uuid"`
}

type User struct {
	ID      int64  `json:"id,omitempty"`
	UUID    string `json:"uuid,omitempty"`
	Code    string `json:"code,omitempty"`
	Name    string `json:"name,omitempty"`
	Phone   string `json:"phone,omitempty"`
	Email   string `json:"email,omitempty"`
	Address string `json:"address,omitempty"`
	Manager *User  `json:"manager,omitempty"`
}

type Base struct {
	UUID      string `json:"uuid"`
	Creator   *User  `json:"creator,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

type Project struct {
	ID        int64            `json:"id,omitempty"`
	Code      string           `json:"code,omitempty"`
	Manager   *User            `json:"manager,omitempty"`
	Category  *ProjectCategory `json:"category,omitempty"`
	Region    *Region          `json:"region,omitempty"`
	Name      string           `json:"name,omitempty"`
	Address   string           `json:"address,omitempty"`
	StartDate string           `json:"start_date,omitempty"`
	EndDate   string           `json:"end_date,omitempty"`
	Tasks     []*Task          `json:"tasks,omitempty"`
	Phases    []*ProjectPhase  `json:"phases,omitempty"`
	Areas     []*ProjectArea   `json:"areas,omitempty"`
}

type ProjectCategory struct {
	ID   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type ProjectPhase struct {
	ID        int64    `json:"id,omitempty"`
	Name      string   `json:"name,omitempty"`
	Project   *Project `json:"project,omitempty"`
	StartDate string   `json:"start_date,omitempty"`
	EndDate   string   `json:"end_date,omitempty"`
}

type ProjectArea struct {
	ID      int64    `json:"id,omitempty"`
	Name    string   `json:"name,omitempty"`
	Project *Project `json:"project,omitempty"`
}

type Task struct {
	ID             int64         `json:"id,omitempty"`
	Code           string        `json:"code,omitempty"`
	Name           string        `json:"name,omitempty"`
	ParentTask     *Task         `json:"parent_task,omitempty"`
	ChildTask      *Task         `json:"child_task,omitempty"`
	ProjectID      int64         `json:"project_id,omitempty"`
	ProjectName    string        `json:"project_name,omitempty"`
	Project        *Project      `json:"project,omitempty"`
	Phase          *ProjectPhase `json:"phase,omitempty"`
	Area           *ProjectArea  `json:"area,omitempty"`
	Creator        *User         `json:"creator,omitempty"`
	Executor       *Executor     `json:"executor,omitempty"`
	Reporter       *User         `json:"reporter,omitempty"`
	Acceptor       *User         `json:"acceptor,omitempty"`
	StartTime      string        `json:"start_time,omitempty"`
	EndTime        string        `json:"end_time,omitempty"`
	Quantity       int64         `json:"quantity,omitempty"`
	Price          int64         `json:"price,omitempty"`
	Unit           string        `json:"unit,omitempty"`
	Description    string        `json:"description,omitempty"`
	Status         string        `json:"status,omitempty"`
	PriorityLevel  string        `json:"priority_level,omitempty"`
	Star           int           `json:"star"`
	TaskImages     []*Image      `json:"task_images,omitempty"`
	OwnerImages    []*Image      `json:"owner_images,omitempty"`
	ExecutorImages []*Image      `json:"executor_images,omitempty"`
	Comments       []*Comment    `json:"comments,omitempty"`
}

type AssigningTask struct {
	Task
	PublicID  int64 `json:"public_id,omitempty"`
	PrivateID int64 `json:"private_id,omitempty"`
	Star      int   `json:"star,omitempty"`
}

type Image struct {
	ID        int64  `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Thumbnail string `json:"thumbnail,omitempty"`
	FullFile  string `json:"full_file,omitempty"`
	Creator   *User  `json:"creator,omitempty"`
}

type Comment struct {
	ID       int64  `json:"id,omitempty"`
	Task     *Task  `json:"task,omitempty"`
	UserID   int64  `json:"user_id,omitempty"`
	UserName string `json:"user_name,omitempty"`
	Creator  *User  `json:"creator,omitempty"`
	Msg      string `json:"msg,omitempty"`
	Time     string `json:"time,omitempty"`
}

type Executor struct {
	ID          int64   `json:"id,omitempty"`
	Code        string  `json:"code,omitempty"`
	Name        string  `json:"name,omitempty"`
	Representer *User   `json:"representer,omitempty"`
	Members     []*User `json:"members,omitempty"`
	Address     string  `json:"address,omitempty"`
	Phone       string  `json:"phone,omitempty"`
	Email       string  `json:"email,omitempty"`
}

type Region struct {
	ID   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
