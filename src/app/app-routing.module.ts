import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { DashboardShopComponent } from './components/dashboard-shop/dashboard-shop.component';
import { UserLoginComponent } from './components/user-login/user-login.component';
import { ProductPageComponent } from './components/dashboard-shop/product-page/product-page.component';
import { PurchasePageComponent } from './components/purchase-page/purchase-page.component';
import { MyshopPageComponent } from './components/myshop-page/myshop-page.component';
import { AddProductPageComponent } from './components/myshop-page/add-product-page/add-product-page.component';
import { EditProductComponent } from './components/myshop-page/edit-product/edit-product.component';
import { OrdersManagePageComponent } from './components/myshop-page/orders-manage-page/orders-manage-page.component';
import { ProfilePageComponent } from './components/profile-page/profile-page.component';
import { HistoryPageComponent } from './components/history-page/history-page.component';
import { PromoCodePageComponent } from './components/myshop-page/promo-code-page/promo-code-page.component';
import { AddPromoCodePageComponent } from './components/myshop-page/promo-code-page/add-promo-code-page/add-promo-code-page.component';

const routes: Routes = [
  { path: '', redirectTo: 'login', pathMatch: 'full' },
  { path: 'login', component: UserLoginComponent },
  { path: 'dashboard', component: DashboardShopComponent },
  { path: 'profile', component: ProfilePageComponent },
  { path: 'dashboard/product/:id', component: ProductPageComponent },
  { path: 'purchase-list/user', component: PurchasePageComponent },
  { path: 'MyShop/user', component: MyshopPageComponent },
  { path: 'MyShop/user/add-product', component: AddProductPageComponent },
  { path: 'MyShop/user/edit-product/:id', component: EditProductComponent },
  { path: 'MyShop/user/orders-manage-page', component: OrdersManagePageComponent },
  { path: 'MyShop/user/promo-code-page', component: PromoCodePageComponent },
  { path: 'MyShop/user/promo-code-page/add', component: AddPromoCodePageComponent },
  { path: 'history', component: HistoryPageComponent}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
