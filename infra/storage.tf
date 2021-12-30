module "storage_accounts" {
  source              = "./modules/storage_accounts"
  resource_group_name = azurerm_resource_group.resource_group.name
  location            = azurerm_resource_group.resource_group.location
  test_prefix         = var.test_prefix
  test_suffix         = var.test_suffix
  subnet_id           = module.networks.subnet_id
}
