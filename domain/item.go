package domain

type Company struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Component struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Item struct {
	ID        int       `json:"id"`
	Company   Company   `json:"company"`
	Component Component `json:"component"`
	Cost      float32   `json:"cost"`
	Price     float32   `json:"price"`
	Comission float32   `json:"comission"`
}

type Machine struct {
	ID         int     `json:"id"`
	Company    Company `json:"company"`
	Components []Item  `json:"components"`
}

type HourPrice struct {
	ID             int     `json:"id"`
	DeploymentTime float32 `json:deploymentTime`
	MonthlyHour    float32 `json:"monthlyHour"`
}

type ServiceOrder struct {
	ID                int       `json:"id"`
	Name              string    `json:"name"`
	HourPrice         HourPrice `json:"hourPrice"`
	Machines          []Machine `json:"machines"`
	DeploymentHour    int       `json:deploymentHour`
	MonthlyHour       int       `json:"monthlyHour"`
	TotalDeployment   float32   `json:"totalDeployment"`
	TotalMonthly      float32   `json:"totalMonthly"`
	TotalServiceOrder float32   `json:"totalServiceOrder"`
}

func NewCompany(company *Company) *Erro {
	if company.Name == "" {
		return &Erro{Codigo: "COMPANY_ERR", Mensagem: "Erro ao cadastrar compania, nome vazio."}
	}
	return nil
}

func NewComponet(component *Component) *Erro {
	if component.Name == "" {
		return &Erro{Codigo: "COMPONET_ERR", Mensagem: "Erro ao cadastrar componet, nome vazio."}
	}
	return nil
}

func NewItem(item *Item) *Erro {
	if item.Component.ID == 0 {
		return &Erro{Codigo: "COMPONET_ERR", Mensagem: "Erro ao cadastrar item, component não associado."}
	}
	if item.Company.ID == 0 {
		return &Erro{Codigo: "COMPONET_ERR", Mensagem: "Erro ao cadastrar item, compania não associada."}
	}
	if item.Cost <= 0 {
		return &Erro{Codigo: "COMPONET_ERR", Mensagem: "Erro ao cadastrar item, custo inválido."}
	}
	if item.Comission < 0 || item.Comission > 100 {
		return &Erro{Codigo: "COMPONET_ERR", Mensagem: "Erro ao cadastrar item, valor de comissão inválida."}
	}
	item.Price = item.Cost * (1 + (item.Comission / 100))
	return nil
}

func NewHourPrice(hourPrice *HourPrice) *Erro {
	if hourPrice.DeploymentTime <= 0 {
		return &Erro{Codigo: "HORU_PRICE_ERR", Mensagem: "Valor da hora de implatação com valor inválido."}
	}
	if hourPrice.MonthlyHour <= 0 {
		return &Erro{Codigo: "HORU_PRICE_ERR", Mensagem: "Valor da hora mensal com valor inválido."}
	}
	return nil
}

func NewMachine(machine *Machine) *Erro {
	if machine.Company.ID == 0 {
		return &Erro{Codigo: "MACHINE_ERR", Mensagem: "Erro ao cadastrar máquina, compania não associada."}
	}
	if len(machine.Components) == 0 {
		return &Erro{Codigo: "MACHINE_ERR", Mensagem: "Erro ao cadastrar máquina, nenhum componente associado."}
	}
	for i := range machine.Components {
		if machine.Components[i].ID == 0 {
			return &Erro{Codigo: "MACHINE_ERR", Mensagem: "Erro ao cadastrar máquina, nenhum component associado."}
		}
	}
	return nil
}

func NewServiceOrder(serviceOrder *ServiceOrder) *Erro {
	if serviceOrder.Name == "" {
		return &Erro{Codigo: "SERVICE_ORDER_ERR", Mensagem: "Erro ao cadastar ordem de serviço, nome vazio."}
	}
	if serviceOrder.HourPrice.ID == 0 {
		return &Erro{Codigo: "SERVICE_ORDER_ERR", Mensagem: "Erro ao cadastrar ordem de serviço, hora não associada."}
	}
	if len(serviceOrder.Machines) == 0 {
		return &Erro{Codigo: "SERVICE_ORDER_ERR", Mensagem: "Erro ao cadastar ordem de servico, nenhuma máquina associada."}
	}
	for i := range serviceOrder.Machines {
		if serviceOrder.Machines[i].ID == 0 {
			return &Erro{Codigo: "MACHINE_ERR", Mensagem: "Erro ao cadastrar ordem de serviço, máquina não associada."}
		}
	}
	if serviceOrder.DeploymentHour <= 0 {
		return &Erro{Codigo: "MACHINE_ERR", Mensagem: "Erro ao cadastrar ordem de serviço, valor inválido para preço de implantação."}
	}
	if serviceOrder.MonthlyHour <= 0 {
		return &Erro{Codigo: "MACHINE_ERR", Mensagem: "Erro ao cadastrar ordem de serviço, valor inválido para preço de manutenção mensal."}
	}
	serviceOrder.TotalDeployment = float32(serviceOrder.DeploymentHour) * serviceOrder.HourPrice.DeploymentTime
	serviceOrder.TotalMonthly = float32(serviceOrder.MonthlyHour) * serviceOrder.HourPrice.MonthlyHour
	serviceOrder.TotalServiceOrder = serviceOrder.TotalDeployment + serviceOrder.TotalMonthly
	for i := range serviceOrder.Machines {
		for j := range serviceOrder.Machines[i].Components {
			cmp := serviceOrder.Machines[i].Components[j]
			serviceOrder.TotalServiceOrder += cmp.Price
		}
	}

	return nil
}
