import { ComponentFixture, TestBed } from '@angular/core/testing';

import { EditProductImageComponent } from './edit-product-image.component';

describe('EditProductImageComponent', () => {
  let component: EditProductImageComponent;
  let fixture: ComponentFixture<EditProductImageComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [EditProductImageComponent]
    });
    fixture = TestBed.createComponent(EditProductImageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
