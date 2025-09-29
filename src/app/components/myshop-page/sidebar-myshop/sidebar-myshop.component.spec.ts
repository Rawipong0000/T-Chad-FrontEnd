import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SidebarMyshopComponent } from './sidebar-myshop.component';

describe('SidebarMyshopComponent', () => {
  let component: SidebarMyshopComponent;
  let fixture: ComponentFixture<SidebarMyshopComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [SidebarMyshopComponent]
    });
    fixture = TestBed.createComponent(SidebarMyshopComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
